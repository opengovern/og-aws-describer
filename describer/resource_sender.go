package describer

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/kaytu-io/kaytu-util/proto/src/golang"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/metadata"
)

const (
	MinBufferSize   int           = 10
	MaxBufferSize   int           = 100
	ChannelSize     int           = 1000
	BufferEmptyRate time.Duration = 5 * time.Second
)

type ResourceSender struct {
	authToken        string
	workspaceName    string
	logger           *zap.Logger
	resourceChannel  chan *golang.AWSResource
	resourceIDs      []string
	doneChannel      chan interface{}
	conn             *grpc.ClientConn
	describeEndpoint string
	jobID            uint
	kafkaTopic       string
	client           golang.DescribeServiceClient

	sendBuffer []*golang.AWSResource
}

func NewResourceSender(workspaceName string, describeEndpoint string, describeToken string, jobID uint, kafkaTopic string, logger *zap.Logger) (*ResourceSender, error) {
	rs := ResourceSender{
		authToken:        describeToken,
		workspaceName:    workspaceName,
		logger:           logger,
		resourceChannel:  make(chan *golang.AWSResource, ChannelSize),
		resourceIDs:      nil,
		doneChannel:      make(chan interface{}),
		conn:             nil,
		describeEndpoint: describeEndpoint,
		kafkaTopic:       kafkaTopic,
		jobID:            jobID,
	}
	if err := rs.Connect(); err != nil {
		return nil, err
	}

	go rs.ResourceHandler()
	return &rs, nil
}

func (s *ResourceSender) Connect() error {
	conn, err := grpc.Dial(
		s.describeEndpoint,
		grpc.WithTransportCredentials(credentials.NewTLS(nil)),
		grpc.WithPerRPCCredentials(oauth.TokenSource{
			TokenSource: oauth2.StaticTokenSource(&oauth2.Token{
				AccessToken: s.authToken,
			}),
		}),
	)
	if err != nil {
		return err
	}
	s.conn = conn

	client := golang.NewDescribeServiceClient(conn)
	s.client = client
	return nil
}

func (s *ResourceSender) ResourceHandler() {
	t := time.NewTicker(BufferEmptyRate)
	defer t.Stop()

	for {
		select {
		case resource := <-s.resourceChannel:
			if resource == nil {
				s.flushBuffer(true)
				s.doneChannel <- struct{}{}
				return
			}

			s.resourceIDs = append(s.resourceIDs, resource.UniqueId)
			s.sendBuffer = append(s.sendBuffer, resource)

			if len(s.sendBuffer) > MaxBufferSize {
				s.flushBuffer(true)
			}
		case <-t.C:
			s.flushBuffer(false)
		}
	}
}

func (s *ResourceSender) flushBuffer(force bool) {
	if len(s.sendBuffer) == 0 {
		return
	}

	if !force && len(s.sendBuffer) < MinBufferSize {
		return
	}

	grpcCtx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
		"workspace-name":  s.workspaceName,
		"resource-job-id": fmt.Sprintf("%d", s.jobID),
	}))

	_, err := s.client.DeliverAWSResources(grpcCtx, &golang.AWSResources{Resources: s.sendBuffer, KafkaTopic: s.kafkaTopic})
	if err != nil {
		s.logger.Error("failed to send resource", zap.Error(err))
		if errors.Is(err, io.EOF) {
			err = s.Connect()
			if err != nil {
				s.logger.Error("failed to reconnect", zap.Error(err))
			}
		}
		return
	}
	s.sendBuffer = nil
}

func (s *ResourceSender) Finish() {
	s.resourceChannel <- nil
	_ = <-s.doneChannel
	s.conn.Close()
}

func (s *ResourceSender) GetResourceIDs() []string {
	return s.resourceIDs
}

func (s *ResourceSender) Send(resource *golang.AWSResource) {
	s.resourceChannel <- resource
}
