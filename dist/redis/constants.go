package console
const consoleTemplate ="<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Redis Cloud Console</title>\n    <script src=\"https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js\"></script>\n\n    <script>\n\n        function load() {\n            let $command = $(\"#command-input\");\n            $command.on(\"keydown\", submitCommand);\n            $command.focus();\n        }\n\n        function submitCommand(e) {\n            let code = (e.keyCode ? e.keyCode : e.which);\n            if (code !== 13) return\n            $.post('redis', {'command': e.data.replace(/\\s+/, ' ')})\n                .done(function (data, status) {\n                    if(status == '200'){\n                        let history = $(\"#history\");\n                        history.empty();\n                        data['history'].forEach(el =>{\n                            history.append('<div>' + data['prompt'] +'&gt;'+ el.input+ '</div>');\n                            history.append('<div>' + el.output+ '</div>');\n\n                        });\n                        let prompt = $(\"#prompt\");\n                        prompt.empty();\n                        prompt.text(data['prompt'] + '>');\n                    }\n\n\n                })\n                .fail(function (xhr, status, error) {\n                    let snackbar = $(\"#snackbar\");\n                    snackbar.text(error);\n                    snackbar.addClass('show');\n                    console.log(error);\n                    setTimeout(function () {\n                        snackbar.removeClass('show')\n                    }, 3000);\n\n                });\n\n        }\n    </script>\n    <style>\n        .console-container {\n            background: black;\n            color: white;\n            height: 100%;\n            width: 100%;\n            display: flex;\n            flex-direction: column;\n            font-family: Monaco, \"Lucida Console\", \"Lucida Grande\", monospace;\n        }\n\n        .console-history {\n            height: 75%;\n\n        }\n\n        .console-input {\n            position: relative;\n            width: 100%;\n        }\n\n        #command-input {\n            width: 75%;\n            background: black;\n            color: white;\n            /*font-size: 1.25em;*/\n            border: none;\n            outline: none;\n            margin-top: 5px;\n            font-family: Monaco, \"Lucida Console\", \"Lucida Grande\", monospace;\n\n        }\n\n        /*input {*/\n        /*    width: 75%;*/\n        /*}*/\n        #prompt {\n            width: 25%;\n            margin-right: 15px;\n        }\n\n        #snackbar {\n            visibility: hidden;\n            min-width: 100%; /* Set a default minimum width */\n            background-color: white; /* Black background color */\n            color: darkred; /* White text color */\n            text-align: center; /* Centered text */\n            border-radius: 2px; /* Rounded borders */\n            padding: 16px; /* Padding */\n            position: fixed; /* Sit on top of the screen */\n            z-index: 1; /* Add a z-index if needed */\n            left: 0; /* Center the snackbar */\n            top: 30px; /* 30px from the bottom */\n        }\n\n        #snackbar.show {\n            visibility: visible; /* Show the snackbar */\n            /* Add animation: Take 0.5 seconds to fade in and out the snackbar.\n            However, delay the fade out process for 2.5 seconds */\n            -webkit-animation: fadein 0.5s, fadeout 0.5s 2.5s;\n            animation: fadein 0.5s, fadeout 0.5s 2.5s;\n        }\n\n        @-webkit-keyframes fadein {\n            from {\n                bottom: 0;\n                opacity: 0;\n            }\n            to {\n                bottom: 30px;\n                opacity: 1;\n            }\n        }\n\n        @keyframes fadein {\n            from {\n                bottom: 0;\n                opacity: 0;\n            }\n            to {\n                bottom: 30px;\n                opacity: 1;\n            }\n        }\n\n        @-webkit-keyframes fadeout {\n            from {\n                bottom: 30px;\n                opacity: 1;\n            }\n            to {\n                bottom: 0;\n                opacity: 0;\n            }\n        }\n\n        @keyframes fadeout {\n            from {\n                bottom: 30px;\n                opacity: 1;\n            }\n            to {\n                bottom: 0;\n                opacity: 0;\n            }\n        }\n    </style>\n</head>\n<body class=\"console-container\" onload=\"load()\">\n<div id=\"snackbar\"></div>\n<div id=\"history\" class=\"console-history\">\n    {{range .history}}\n        <div>{{.prompt}}&gt;&nbsp;{{.input}}</div>\n        <div>{{.output}}</div>\n    {{end}}\n</div>\n{{/*<div id=\"history\" class=\"console-history\">*/}}\n{{/*    <div>history item one</div>*/}}\n{{/*    <div>command output</div>*/}}\n{{/*</div>*/}}\n<div id=\"input\" class=\"console-input\">\n    <span id=\"prompt\">{{.prompt}}&gt;</span><input type=\"text\" id=\"command-input\">\n\n</div>\n\n\n</body>\n</html>"

