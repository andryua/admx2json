//const json = require.open("GET", "/admjson")

$(function() {
    $("#tree").fancytree({
        checkbox: false,
        selectMode: 3,
        extensions: ["filter"],
        quicksearch: false,

        source: $.ajax({
            url: "/admjson",
            cache: true,
            dataType: "json",
        }),

        filter: {
            autoApply: false, // Re-apply last filter if lazy data is loaded
            autoExpand: false, // Expand all branches that contain matches while filtered
            counter: true, // Show a badge with number of matching child nodes near parent icons
            fuzzy: false, // Match single characters in order, e.g. 'fb' will match 'FooBar'
            hideExpandedCounter: true, // Hide counter badge if parent is expanded
            hideExpanders: false, // Hide expanders if all child nodes are hidden by filter
            highlight: true, // Highlight matches by wrapping inside <mark> tags
            leavesOnly: false, // Match end nodes only
            nodata: true, // Display a 'no data' status node if result is empty
            mode: "hide" // Grayout unmatched nodes (pass "hide" to remove unmatched node instead)
        },

        activate: function(event, data) {
            //var reginfo = data.node.data.reginfo.replace(/!/g,"-->")
            //reginfo = reginfo.replace(/ HK/g,"<br>HK")
            if (data.node.title && !data.node.folder) {
                $("#tit").show();
                $("#echoName").text(data.node.title)
                $("#tit1").show();
                $("#tit2").show();
                $("#tit3").show();
                $("#echoSupport").text(data.node.data.support);
                //$("#echoRegInfo").html(reginfo);
//                console.log(node.data.vaues)
                $("#echoInfo").text(data.node.data.description);
                $("#echoValues").text(JSON.stringify(data.node.data.values));

            } else {

                $("#tit").hide();
                $("#tit1").hide();
                $("#tit2").hide();
                $("#tit3").hide();
                $("#echoName").text("")
                $("#echoSupport").text("");
                $("#echoInfo").text("");
                $("#echoValues").text("");
                //$("#echoRegInfo").text("");
            };
        },
        deactivate: function(event, data) {
            // $("#tit").hide();
            $("#tit1").hide();
            $("#tit2").hide();
            $("#tit3").hide();
            $("#echoSupport").text("");
            $("#echoInfo").text("");
            $("#echoName").text("");
            $("#echoRegInfo").text("");
        },
        select: function(event, data) {
            var selKeys = $.map(tree.getSelectedNodes(), function(node) {
                if (!node.folder)
                    return node.data.id;
            });

            $("#admtmpid").val(function() {
                var emphasis = selKeys.join(",");
                return emphasis;
            });

        },
    });

    var tree = $("#tree").fancytree("getTree");

    $(".fancytree-container").addClass("fancytree-connectors");

    $("input[name=search]").keyup(function(e) {
        var n,
            tree = $.ui.fancytree.getTree(),
            args = "autoApply autoExpand fuzzy hideExpanders highlight leavesOnly nodata".split(" "),
            opts = {},
            filterFunc = tree.filterNodes,
            match = $(this).val();

        opts.mode = "hide";
        //opts.hideExpandedCounter = true;
        //opts.autoExpand = false;

        if (e && e.which === $.ui.keyCode.ESCAPE || $.trim(match) === "") {
            $("button#btnResetSearch").click();
            return;
        }
        n = filterFunc.call(tree, match, opts);
        $("button#btnResetSearch").attr("disabled", false);
        //$("span#matches").text("(" + n + " matches)");
    }).focus();

    $("button#btnResetSearch").click(function(e) {
        $("input[name=search]").val("");
        $("span#matches").text("");
        tree.clearFilter();
    }).attr("disabled", true);


  
    
    $("button#reset").click (function() {
        $.post("/resetdb", function(data) {
            if (data == "done") {
                location.reload(true);
            }
         });
    });

    function FuncNAme(id, name) {
        var uri = "/some/path?id=" + id + "&name=" + name;
        var data = $("#desttree").text()
        $.post(uri, data, function() {

        })
    }

    $("button#next").click(function() {
        var obj = new Object();
        var jsonString;
        var selKeys = $.map(tree.getSelectedNodes(), function(node) {
            if (!node.folder) {
                obj.id = node.data.id;
                obj.title = node.data.title;
                obj.support = node.data.support;
                console.log(node.data.vaues)
                obj.reginfo = node.data.values.toString();
                obj.info = node.data.description;
                jsonString = JSON.stringify(obj);
                return jsonString;
            }
        });
        $("#json").html(function() {
            var emphasis = selKeys.join(",");
            return emphasis;
        });

        /*$.post(
              "/url", { selected: vals },
              function(data) {
                  console.log(data)
              }
          )*/


    });

});

