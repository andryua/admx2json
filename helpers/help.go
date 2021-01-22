package helpers

func createJS (presentation Presentation_json) string {
	script := "<div class=\"form-check form-switch\"><input class=\"form-check-input\" type=\"checkbox\" id=\"flexSwitchCheckDefault\"><label class=\"form-check-label\" for=\"flexSwitchCheckDefault\">Enable or disable rule</label></div>"
	if presentation.Text != nil {
		for _,txt := range presentation.Text {
			script = "<label for=\"exampleDataList\" class=\"form-label\">" + txt + "</label>"
		}
	}
	if presentation.TextBox != nil {
		for _, txtb := range presentation.TextBox {
			script = "<div id=" + presentation.ID + "class=\"mb-3\"><label for=\"exampleDataList\" class=\"form-label\">" + txtb.Label + "</label><input class=\"form-control\" list=\"datalistOptions\" id=" + txtb.RefId + "placeholder=" + txtb.DefaultValue + "></div>"
		}
	}
	if presentation.MultiTextBox != nil {
		for mtxtb, _ := range presentation.MultiTextBox {

		}
	}
	if presentation.LongDecimalTextBox != nil {
		for ldtxtb, _ := range presentation.LongDecimalTextBox {

		}
	}
	if presentation.ListBox != nil {
		for lb, _ := range presentation.ListBox {

		}
	}
	if presentation.DropdownList != nil {
		for _, ddl := range presentation.DropdownList {
			script = "<select id=" + ddl.RefId + "class=\"form-select\" aria-label=\"Default select example\">\n  <option selected>" + ddl.DefaultItem + "</option>\n  <option value=\"1\">One</option>\n  <option value=\"2\">Two</option>\n  <option value=\"3\">Three</option>\n</select>"
		}
	}if presentation.DecimalTextBox != nil {
		for dtxtb, _ := range presentation.DecimalTextBox {

		}
	}
	if presentation.ComboBox != nil {
		for combo, _ := range presentation.ComboBox {

		}
	}
	if presentation.CheckBox != nil {
		for chkb, _ := range presentation.CheckBox {

		}
	}

	return script

}