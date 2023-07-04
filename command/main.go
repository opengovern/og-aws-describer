/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/kaytu-io/kaytu-aws-describer/aws"
	"github.com/kaytu-io/kaytu-util/pkg/describe/enums"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	checkAttachedPolicies                         bool
	resourceType, accessKey, accountID, secretKey string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kaytu-aws-describer",
	Short: "kaytu aws describer manual",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger, _ := zap.NewProduction()
		cfg, err := aws.GetConfig(context.Background(), accessKey, secretKey, "", "", nil)
		if checkAttachedPolicies {
			isAttached, err := aws.CheckAttachedPolicy(logger, cfg, aws.SecurityAuditPolicyARN)
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
	rootCmd.Flags().BoolVarP(&checkAttachedPolicies, "checkAttachedPolicies", "", false, "Check attached policies")
	rootCmd.Flags().StringVarP(&resourceType, "resourceType", "t", "", "Resource type")
	rootCmd.Flags().StringVarP(&accountID, "accountID", "", "", "AccountID")
	rootCmd.Flags().StringVarP(&accessKey, "accessKey", "a", "", "Access key")
	rootCmd.Flags().StringVarP(&secretKey, "secretKey", "s", "", "Secret key")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
