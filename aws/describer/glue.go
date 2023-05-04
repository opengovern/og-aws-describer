package describer

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func GlueCatalogDatabase(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := glue.NewFromConfig(cfg)
	paginator := glue.NewGetDatabasesPaginator(client, &glue.GetDatabasesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, database := range page.DatabaseList {
			arn := fmt.Sprintf("arn:aws:glue:%s:%s:database/%s", describeCtx.Region, describeCtx.AccountID, *database.Name)
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				Name:   *database.Name,
				ARN:    arn,
				Description: model.GlueCatalogDatabaseDescription{
					Database: database,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func GlueCatalogTable(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := glue.NewFromConfig(cfg)
	paginator := glue.NewGetDatabasesPaginator(client, &glue.GetDatabasesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, database := range page.DatabaseList {
			tablePaginator := glue.NewGetTablesPaginator(client, &glue.GetTablesInput{
				DatabaseName: database.Name,
				CatalogId:    database.CatalogId,
			})
			for tablePaginator.HasMorePages() {
				tablePage, err := tablePaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				for _, table := range tablePage.TableList {
					arn := fmt.Sprintf("arn:aws:glue:%s:%s:table/%s/%s", describeCtx.Region, describeCtx.AccountID, *database.Name, *table.Name)
					resource := Resource{
						Region: describeCtx.KaytuRegion,
						Name:   *table.Name,
						ARN:    arn,
						Description: model.GlueCatalogTableDescription{
							Table: table,
						},
					}
					if stream != nil {
						if err := (*stream)(resource); err != nil {
							return nil, err
						}
					} else {
						values = append(values, resource)
					}
				}
			}
		}
	}

	return values, nil
}

func GlueConnection(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := glue.NewFromConfig(cfg)
	paginator := glue.NewGetConnectionsPaginator(client, &glue.GetConnectionsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, connection := range page.ConnectionList {
			arn := fmt.Sprintf("arn:aws:glue:%s:%s:connection/%s", describeCtx.Region, describeCtx.AccountID, *connection.Name)
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				Name:   *connection.Name,
				ARN:    arn,
				Description: model.GlueConnectionDescription{
					Connection: connection,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func GlueCrawler(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := glue.NewFromConfig(cfg)
	paginator := glue.NewGetCrawlersPaginator(client, &glue.GetCrawlersInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, crawler := range page.Crawlers {
			arn := fmt.Sprintf("arn:aws:glue:%s:%s:crawler/%s", describeCtx.Region, describeCtx.AccountID, *crawler.Name)
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				Name:   *crawler.Name,
				ARN:    arn,
				Description: model.GlueCrawlerDescription{
					Crawler: crawler,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func GetGlueCrawler(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	name := fields["name"]

	client := glue.NewFromConfig(cfg)

	out, err := client.GetCrawler(ctx, &glue.GetCrawlerInput{
		Name: &name,
	})
	if err != nil {
		return nil, err
	}
	crawler := out.Crawler

	var values []Resource
	arn := fmt.Sprintf("arn:aws:glue:%s:%s:crawler/%s", describeCtx.Region, describeCtx.AccountID, *crawler.Name)
	values = append(values, Resource{
		Region: describeCtx.KaytuRegion,
		Name:   *crawler.Name,
		ARN:    arn,
		Description: model.GlueCrawlerDescription{
			Crawler: *crawler,
		},
	})

	return values, nil
}

func GlueDataCatalogEncryptionSettings(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := glue.NewFromConfig(cfg)

	settings, err := client.GetDataCatalogEncryptionSettings(ctx, &glue.GetDataCatalogEncryptionSettingsInput{})
	if err != nil {
		return nil, err
	}

	var values []Resource
	resource := Resource{
		Region: describeCtx.KaytuRegion,
		Description: model.GlueDataCatalogEncryptionSettingsDescription{
			DataCatalogEncryptionSettings: *settings.DataCatalogEncryptionSettings,
		},
	}
	if stream != nil {
		if err := (*stream)(resource); err != nil {
			return nil, err
		}
	} else {
		values = append(values, resource)
	}

	return values, nil
}

func GlueDataQualityRuleset(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	//
	client := glue.NewFromConfig(cfg)
	paginator := glue.NewListDataQualityRulesetsPaginator(client, &glue.ListDataQualityRulesetsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, listRuleset := range page.Rulesets {
			ruleset, err := client.GetDataQualityRuleset(ctx, &glue.GetDataQualityRulesetInput{
				Name: listRuleset.Name,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				Name:   *listRuleset.Name,
				Description: model.GlueDataQualityRulesetDescription{
					DataQualityRuleset: *ruleset,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func GlueDevEndpoint(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := glue.NewFromConfig(cfg)
	paginator := glue.NewGetDevEndpointsPaginator(client, &glue.GetDevEndpointsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, devEndpoint := range page.DevEndpoints {
			arn := fmt.Sprintf("arn:aws:glue:%s:%s:devEndpoint/%s", describeCtx.Region, describeCtx.AccountID, *devEndpoint.EndpointName)
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				Name:   *devEndpoint.EndpointName,
				ARN:    arn,
				Description: model.GlueDevEndpointDescription{
					DevEndpoint: devEndpoint,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func GlueJob(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := glue.NewFromConfig(cfg)
	paginator := glue.NewGetJobsPaginator(client, &glue.GetJobsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, job := range page.Jobs {
			arn := fmt.Sprintf("arn:aws:glue:%s:%s:job/%s", describeCtx.Region, describeCtx.AccountID, *job.Name)

			bookmark, err := client.GetJobBookmark(ctx, &glue.GetJobBookmarkInput{
				JobName: job.Name,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				Name:   *job.Name,
				ARN:    arn,
				Description: model.GlueJobDescription{
					Job:      job,
					Bookmark: *bookmark.JobBookmarkEntry,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func GetGlueJob(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	jobName := fields["name"]

	client := glue.NewFromConfig(cfg)

	var values []Resource
	out, err := client.GetJob(ctx, &glue.GetJobInput{
		JobName: &jobName,
	})
	if err != nil {
		return nil, err
	}
	job := out.Job

	arn := fmt.Sprintf("arn:aws:glue:%s:%s:job/%s", describeCtx.Region, describeCtx.AccountID, *job.Name)

	bookmark, err := client.GetJobBookmark(ctx, &glue.GetJobBookmarkInput{
		JobName: job.Name,
	})
	if err != nil {
		return nil, err
	}

	values = append(values, Resource{
		Region: describeCtx.KaytuRegion,
		Name:   *job.Name,
		ARN:    arn,
		Description: model.GlueJobDescription{
			Job:      *job,
			Bookmark: *bookmark.JobBookmarkEntry,
		},
	})

	return values, nil
}

func GlueSecurityConfiguration(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := glue.NewFromConfig(cfg)
	paginator := glue.NewGetSecurityConfigurationsPaginator(client, &glue.GetSecurityConfigurationsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, securityConfiguration := range page.SecurityConfigurations {
			arn := fmt.Sprintf("arn:aws:glue:%s:%s:security-configuration/%s", describeCtx.Region, describeCtx.AccountID, *securityConfiguration.Name)
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				Name:   *securityConfiguration.Name,
				ARN:    arn,
				Description: model.GlueSecurityConfigurationDescription{
					SecurityConfiguration: securityConfiguration,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}
