package sample

test_design_group_kpi_editor {
	allow with input as {"issuer": "service-a", "audience": "service-b"}
	not allow with input as {"issuer": "service-a", "audience": "service-b"}
}
