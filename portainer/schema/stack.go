package schema

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func StacksSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"stacks": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"additional_files": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"auto_update": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"interval": schema.StringAttribute{Computed: true},
								"job_id":   schema.StringAttribute{Computed: true},
								"webhook":  schema.StringAttribute{Computed: true},
							},
						},
						"endpoint_id": schema.Int64Attribute{
							Computed: true,
						},
						"entry_point": schema.StringAttribute{
							Computed: true,
						},
						"env": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Computed: true,
									},
									"value": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"id": schema.Int64Attribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"option": schema.ObjectAttribute{
							Computed: true,
							AttributeTypes: map[string]attr.Type{
								"prune": types.BoolType,
							},
						},
						"resource_control": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"access_level":        schema.Int64Attribute{Computed: true},
								"administrators_only": schema.BoolAttribute{Computed: true},
								"id":                  schema.Int64Attribute{Computed: true},
								"owner_id":            schema.Int64Attribute{Computed: true},
								"public":              schema.BoolAttribute{Computed: true},
								"resource_id":         schema.StringAttribute{Computed: true},
								"sub_resource_ids": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
								"system": schema.BoolAttribute{Computed: true},
								"team_accesses": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"access_level": schema.Int64Attribute{Computed: true},
											"team_id":      schema.Int64Attribute{Computed: true},
										},
									},
								},
								"type": schema.Int64Attribute{Computed: true},
								"user_accesses": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"access_level": schema.Int64Attribute{Computed: true},
											"user_id":      schema.Int64Attribute{Computed: true},
										},
									},
								},
							},
						},
						"status": schema.Int64Attribute{
							Computed: true,
						},
						"swarm_id": schema.StringAttribute{
							Computed: true,
						},
						"type": schema.Int64Attribute{
							Computed: true,
						},
						"created_by": schema.StringAttribute{
							Computed: true,
						},
						"creation_date": schema.Int64Attribute{
							Computed: true,
						},
						"from_app_template": schema.BoolAttribute{
							Computed: true,
						},
						"git_config": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"authentication": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"git_credential_id": schema.Int64Attribute{Computed: true},
										"password": schema.StringAttribute{
											Computed: true,
										},
										"username": schema.StringAttribute{Computed: true},
									},
								},
								"config_file_path": schema.StringAttribute{Computed: true},
								"config_hash":      schema.StringAttribute{Computed: true},
								"reference_name":   schema.StringAttribute{Computed: true},
								"url":              schema.StringAttribute{Computed: true},
							},
						},
						"is_compose_format": schema.BoolAttribute{
							Computed: true,
						},
						"namespace": schema.StringAttribute{
							Computed: true,
						},
						"project_path": schema.StringAttribute{
							Computed: true,
						},
						"update_date": schema.Int64Attribute{
							Computed: true,
						},
						"updated_by": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}
