function funkCheck(f, gplist, event) {
    var gp = gplist.split(",");
    var idv = document.getElementById("id").value;
    var idchk;
    if (idv != "") {
        idchk = document.getElementById(idv).value;
    }
    var nm = document.getElementById("gpname").value;
    //console.log(gp);

    //console.log(nm);
    var chck1 = document.getElementById("recipient-name");
    var chck2 = document.getElementById("recipient-name22");
    var btn1 = document.getElementById("addgp");
    var btn2 = document.getElementById("editgp");

    if (event == "add") {
        if (gp.includes(chck1.value)) {
            btn1.disabled = true;
            document.getElementById("checkAdd").innerHTML = "<div class='alert alert-danger alert-dismissible fade show' role='alert'>Таке ім'я вже існує! Введіть інше, будь ласка.<button type='button' class='close' data-dismiss='alert' aria-label='Close' id='enableAdd' onclick='funcBtnEnabl(this)'><span aria-hidden='true'>&times;</span></button></div>";
        } else {
            btn1.disabled = false;
            f.submit();
        };
    } else {
        if (chck2.value != nm) {
            if (gp.includes(chck2.value)) {
                btn2.disabled = true;
                document.getElementById("checkEdit").innerHTML = "<div class='alert alert-danger alert-dismissible fade show' role='alert'>Tаке ім'я вже існує! Введіть інше, будь ласка.<button type='button' class='close' data-dismiss='alert' aria-label='Close' id='enableEdit' onclick='funcBtnEnabl(this)'><span aria-hidden='true'>&times;</span></button></div>";
            } else {
                btn2.disabled = false;
                f.submit();
            };
        } else {
            btn2.disabled = false;
            f.submit();
        };
    };

};

function funcBtnEnabl(id) {
    //console.log(id.id);
    var btn1 = document.getElementById('addgp');
    var btn2 = document.getElementById('editgp');
    if (id.id == "enableAdd") {
        btn1.disabled = false;
    } else {
        btn2.disabled = false;
    };
};

function funcEditModal(id, name, info, depend, unchange) {
    var idv = document.getElementById("id");
    var nm = document.getElementById("gpname");
    var uch = document.getElementById("unchange");
    var nam = document.getElementById("recipient-name22");
    var inf = document.getElementById("message-text22");
    var dep = document.getElementById("chckbx");
    var hide = document.getElementById("hide");
    var collapse = document.getElementById("collapse22");
    var colledit = document.getElementById("collapsedit");
    for (i = 0; i < dep.getElementsByTagName("input").length; i++) {
        dep.getElementsByTagName("input")[i].checked = false;
    };
    //nam.placeholder = name;
    //inf.placeholder = info;
    nam.value = name;
    inf.value = info;
    var depen = depend.split(",");
    if (depen == "") {
        for (i = 0; i < dep.getElementsByTagName("input").length; i++) {
            dep.getElementsByTagName("input")[i].checked = false;
        }
    } else {
        // console.log(depend);
        // console.log(depen);
        //console.log(dep.getElementsByTagName("input"));
        for (i = 0; i < dep.getElementsByTagName("input").length; i++) {
            for (j = 0; j < depen.length; j++) {
                //      console.log("values %s", depen[j]);

                if (dep.getElementsByTagName("input")[i].value == depen[j]) {
                    dep.getElementsByTagName("input")[i].checked = true;
                }
            }
        }
    };
    idv.value = id;
    nm.value = name;
    uch.value = unchange;
    if (unchange == "1") {
        //hide.style.visibility = "hidden";
        //collapse.class = "collapse";
        //colledit.class = "btn btn-light collapsed";
        //colledit.setAttribute("aria-expanded", false);
        //hide.disabled = true;
        for (i = 0; i < dep.getElementsByTagName("input").length; i++) {
            dep.getElementsByTagName("input")[i].disabled = true;
        }
    } else {
        for (i = 0; i < dep.getElementsByTagName("input").length; i++) {
            dep.getElementsByTagName("input")[i].disabled = false;
        }
        //hide.style.visibility = 'visible';
        //collapse.class = "collapse show";
        //colledit.class = "btn btn-light";
        //colledit.setAttribute("aria-expanded", true);
        //hide.disabled = false;
    };
    /*if (unchange == "0") {
        //hide.style.visibility = 'visible';
        collapse.class = "collapse show";
        colledit.class = "btn btn-light";
        colledit.setAttribute("aria-expanded", true);
        //hide.disabled = false;
    } else {
        //hide.style.visibility = "hidden";
        collapse.class = "collapse";
        colledit.class = "btn btn-light collapsed";
        colledit.setAttribute("aria-expanded", false);
        //hide.disabled = true;
    };*/
};

//функція для деактивації поля при роботі з значеннями реєстру
function func1(radio, regval, valreg) {
    var a = document.getElementById(valreg);
    var b = document.getElementById(regval);
    if (radio.checked) {
        a.disabled = true;
        b.disabled = false;
    }
};
//при виборі видалення - деактивується поле для введення значення
function valDisable(select, value) {
    var b = document.getElementById(value);
    if (select.value == "DELETEALLVALUES" || one.value == "DELETE") {
        b.disabled = true;
    } else {
        b.disabled = false;
    };
};
//функція збереження змін в правило - не перевантажує всю стрінку
function fSave(id, gpid) {
    //var token = document.getElementById("token").value;
    $.post(
        "/update?id=" + id + "&gpid=" + gpid,// + "&token=" + token,
        $("#update-form-" + id).serialize(),
        function(data) {
            $("#res" + id).html("<div class='alert alert-success alert-dismissible fade show' role='alert'>Зміни збережено!<button type='button' class='close' data-dismiss='alert' aria-label='Close'><span aria-hidden='true'>&times;</span></button></div>");
        },
    );
};
//функція для блокування редагування правил з типових ГП
function funcDisabled(upd, gpid, gpname, id) {
    var dep = document.getElementById("depend");
    var gpid = document.getElementById(gpid);
    var name = document.getElementById("gpglobname");
    var gpname = document.getElementById(gpname);
    var gpgid = document.getElementById("gpglobid");
    var cp = document.getElementById("copygpo-" + id);
    //var token = document.getElementById("token").value
    cp.href = "/copy?id=" + id + "&name=" + name.value + "&gpid=" + gpgid.value;// + "&token=" + token;
    //console.log(cp.href)
    depend = (dep.value).split(",");
    if (gpid.value != gpgid.value) {
        //        $("#" + one.id + " :input, #" + one.id + " :select, #" + one.id + " :radio").attr("disabled", true);
        //$("#update-form-{{.Id}} input, #update-form-{{.Id}} select").attr('disabled', true);
        var elements = document.getElementById(upd).elements;
        for (var i = 0, len = elements.length; i < len; ++i) {
            elements[i].disabled = true;
            //console.log(elements[i].name);
        }
    }
};


function funkCheckEdit(f, gplist, txt) {
    var gp = gplist.split(",");

    var nm = document.getElementById("gpglobname").value;
    //console.log(gp);

    //console.log(nm);
    var chck = document.getElementById("recipient-name22");
    var btn = document.getElementById("editgp");

    if (chck.value != nm) {
        if (gp.includes(chck.value)) {
            btn.disabled = true;
            document.getElementById("checkEdit").innerHTML = "<div class='alert alert-danger alert-dismissible fade show' role='alert'>Tаке ім'я вже існує! Введіть інше, будь ласка.<button type='button' class='close' data-dismiss='alert' aria-label='Close' id='enableEdit' onclick='funcBtnEnablEdit(this)'><span aria-hidden='true'>&times;</span></button></div>";
        } else {
            btn.disabled = false;
            f.submit();
        };
    } else {
        btn.disabled = false;
        f.submit();
    };


};

function funcBtnEnablEdit(id) {
    //console.log(id.id);
    var btn = document.getElementById('editgp');
    if (id.id == "enableEdit") {
        btn.disabled = false;
    };
};

function funcEditModalEdit(id, name, info, depend, unchange) {
    //console.log(id, name, info, depend, unchange);
    var nam = document.getElementById("recipient-name22");
    var inf = document.getElementById("message-text22");
    var dep = document.getElementById("chckbx");
    var idv = document.getElementById("id");
    var nm = document.getElementById("gpname");
    /*for (i = 0; i < dep.getElementsByTagName("input").length; i++) {
        dep.getElementsByTagName("input")[i].checked = false;
    };*/
    nam.value = name;
    inf.value = info;
    idv.value = id;
    nm.value = name;
    var depen = depend.split(",");
    if (unchange == "0") {
        if (depen == "") {
            for (i = 0; i < dep.getElementsByTagName("input").length; i++) {
                dep.getElementsByTagName("input")[i].checked = false;
            }
        } else {

            for (i = 0; i < dep.getElementsByTagName("input").length; i++) {
                for (j = 0; j < depen.length; j++) {

                    if (dep.getElementsByTagName("input")[i].value == depen[j]) {
                        dep.getElementsByTagName("input")[i].checked = true;
                    }
                }
            }
        }
    };


};

function funcSave (button,id) {
    $.post(
        "/export?id=" + id,
        //$("#update-form-" + id).serialize(),
        function(data) {
            obj = JSON.parse(data)
            //console.log (obj);
            if (obj.success) {
                //console.log (obj.data.policy);
                //$("#res-" + id).html("<div class='alert alert-success alert-dismissible fade show' role='alert'>Зміни збережено!<button type='button' class='close' data-dismiss='alert' aria-label='Close'><span aria-hidden='true'>&times;</span></button></div>");
                document.getElementById("res-"+id).style = "font-size:10;color:green";
                $("#res-" + id).hide().html("Зміни збережено!").fadeIn('slow').delay(5000).hide(1);
            } else {
                //console.log (obj);
                //$("#res-" + id).html("<div class='alert alert-danger alert-dismissible fade show' role='alert'>Зміни не збережено!<button type='button' class='close' data-dismiss='alert' aria-label='Close'><span aria-hidden='true'>&times;</span></button></div>");
                document.getElementById("res-"+id).style = "font-size:10;color:red";
                $("#res-" + id).hide().html("Зміни не збережено!").fadeIn('slow').delay(5000).hide(1);
            }
        },
    );
};

function funcHref (button) {
    id = (button.id).split("-")[1]
    button.href = button.href;// + document.getElementById("token").value;
    if (button.id == "asave-"+id) {
        $.post(
            "/export?id=" + id,// + "&token=" + document.getElementById("token").value,
            //$("#update-form-" + id).serialize(),
            function(data) {
                if (data == "saved") {
                    document.getElementById("res-"+id).style = "font-size:10;color:green";
                    $("#res-" + id).hide().html("Зміни збережено!").fadeIn('slow').delay(5000).hide(1);
                }
            })
        }
};


function funcallgp() {
    $(".usr").show();
    $(".def").show();
};
function funcusrgp() {
    $(".usr").show();
    $(".def").hide();
};
function funcdefgp() {
    $(".usr").hide();
    $(".def").show();
};
