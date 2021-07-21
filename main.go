package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
)

var testPolicyTags = &bigquery.PolicyTagList{Names: []string{"projects/your_project/locations/your_location/taxonomies/your_taxonomies/policyTags/your_tag"}}

func main() {

	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, "your_project")
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	defer client.Close()

	dataset := client.Dataset("your_dataset")
	table := dataset.Table("your_table")

	if err = table.Delete(ctx); err != nil {
		fmt.Printf("err: %v", err)
	}

	metaData := &bigquery.TableMetadata{Schema: schema()}
	if err := table.Create(ctx, metaData); err != nil {
		fmt.Printf("err: %v", err)
	}

}

func schema() bigquery.Schema {
	return []*bigquery.FieldSchema{
		{Name: "id", Required: true, Type: bigquery.IntegerFieldType},
		{Name: "username", Type: bigquery.StringFieldType, PolicyTags: testPolicyTags},
		{Name: "email", Type: bigquery.StringFieldType, PolicyTags: testPolicyTags},
		{Name: "phone_number", Type: bigquery.StringFieldType, PolicyTags: testPolicyTags},
		{Name: "created_at", Required: true, Type: bigquery.TimestampFieldType},
		{Name: "updated_at", Required: true, Type: bigquery.TimestampFieldType},
	}
}
