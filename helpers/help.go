package helpers

func createJS(presentation Presentation_json) string {
	script := ""
	script = "<div class=\"form-check form-switch\"><input class=\"form-check-input\" type=\"checkbox\" id=\"flexSwitchCheckDefault\"><label class=\"form-check-label\" for=\"flexSwitchCheckDefault\">Enable or disable rule</label></div>"
	if presentation.Text != nil {
		for _, txt := range presentation.Text {
			script += "<label for=\"exampleDataList\" class=\"form-label\">" + txt + "</label>"
		}
	}
	if presentation.TextBox != nil {
		for _, txtb := range presentation.TextBox {
			script += "<div id=" + presentation.ID + "class=\"mb-3\"><label for=\"exampleDataList\" class=\"form-label\">" + txtb.Label + "</label><input class=\"form-control\" list=\"datalistOptions\" id=" + txtb.RefId + "placeholder=" + txtb.DefaultValue + "></div>"
		}
	}
	if presentation.MultiTextBox != nil {
		for _, mtxtb := range presentation.MultiTextBox {
			script += "<div class=\"mb-3\" id=" + mtxtb.RefId + ">\n  <label for=\"exampleFormControlTextarea1\" class=\"form-label\">Enter values here</label>\n  <textarea class=\"form-control\" id=\"exampleFormControlTextarea1\" rows=\"3\"></textarea>\n</div>"
		}
	}
	if presentation.LongDecimalTextBox != nil {
		for _, ldtxtb := range presentation.LongDecimalTextBox {
			script += "<div id=" + presentation.ID + "class=\"mb-3\"><label for=\"exampleDataList\" class=\"form-label\">" + ldtxtb.Text + "</label><input class=\"form-control\" list=\"datalistOptions\" id=" + ldtxtb.RefId + "placeholder=" + ldtxtb.DefaultValue + "></div>"
		}
	}
	if presentation.ListBox != nil {
		for _, lb := range presentation.ListBox {
			script += "<ul class=\"list-group\" id= " + lb.RefId + ">\n  <li class=\"list-group-item\">Cras justo odio</li>\n  <li class=\"list-group-item\">Dapibus ac facilisis in</li>\n  <li class=\"list-group-item\">Morbi leo risus</li>\n  <li class=\"list-group-item\">Porta ac consectetur ac</li>\n  <li class=\"list-group-item\">Vestibulum at eros</li>\n</ul>"
		}
	}
	if presentation.DropdownList != nil {
		for _, ddl := range presentation.DropdownList {
			script += "<select id=" + ddl.RefId + "class=\"form-select\" aria-label=\"Default select example\">\n  <option selected>" + ddl.DefaultItem + "</option>\n  <option value=\"1\">One</option>\n  <option value=\"2\">Two</option>\n  <option value=\"3\">Three</option>\n</select>"
		}
	}
	if presentation.DecimalTextBox != nil {
		for _, dtxtb := range presentation.DecimalTextBox {
			script += "<div id=" + presentation.ID + "class=\"mb-3\"><label for=\"exampleDataList\" class=\"form-label\">" + dtxtb.Text + "</label><input class=\"form-control\" list=\"datalistOptions\" id=" + dtxtb.RefId + "placeholder=" + dtxtb.DefaultValue + "></div>"
		}
	}
	if presentation.ComboBox != nil {
		for _, combo := range presentation.ComboBox {
			script += "<label for=\"exampleDataList\" class=\"form-label\">" + combo.Text + "</label>"
		}
	}
	if presentation.CheckBox != nil {
		for _, chkb := range presentation.CheckBox {
			script += "<div class=\"form-check\" id=" + chkb.RefId + ">\n  <input class=\"form-check-input\" type=\"checkbox\" value=\"\" id=\"flexCheckDefault\" " + chkb.DefaultChecked + ">\n  <label class=\"form-check-label\" for=\"flexCheckDefault\">\n" + chkb.Text + "\n  </label>\n</div>"
		}
	}

	return script
}
