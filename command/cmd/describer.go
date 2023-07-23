/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kaytu-io/kaytu-aws-describer/aws"
	"github.com/kaytu-io/kaytu-util/pkg/describe/enums"
	"github.com/spf13/cobra"
)

var (
	checkAttachedPolicies                         bool
	resourceType, accessKey, accountID, secretKey string
)

// describerCmd represents the describer command
var describerCmd = &cobra.Command{
	Use:   "describer",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := aws.GetConfig(context.Background(), accessKey, secretKey, "", "", nil)
		if checkAttachedPolicies {
			isAttached, err := aws.CheckAttachedPolicy(cfg, aws.SecurityAuditPolicyARN)
			fmt.Println("IsAttached", isAttached)
			fmt.Println("Error", err)
			return nil
		}

		output, err := aws.GetResources(
			context.Background(),
			resourceType,
			enums.DescribeTriggerTypeManual,
			accountID,
			nil,
			accessKey,
			secretKey,
			"",
			"",
			nil,
			false,
			nil,
		)
		if err != nil {
			return fmt.Errorf("AWS: %w", err)
		}
		js, err := json.Marshal(output)
		if err != nil {
			return err
		}
		fmt.Println(string(js))
		return nil
	},
}

func init() {
	describerCmd.Flags().BoolVar(&checkAttachedPolicies, "checkAttachedPolicies", false, "Check attached policies")
	describerCmd.Flags().StringVar(&resourceType, "resourceType", "", "Resource type")
	describerCmd.Flags().StringVar(&accountID, "accountID", "", "AccountID")
	describerCmd.Flags().StringVar(&accessKey, "accessKey", "", "Access key")
	describerCmd.Flags().StringVar(&secretKey, "secretKey", "", "Secret key")
}
