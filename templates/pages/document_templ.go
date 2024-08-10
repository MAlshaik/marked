// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/MAlshaik/marked/models"
import "github.com/MAlshaik/marked/templates/layouts"

func DocumentPage(document models.Document) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"mx-auto\"><div class=\"flex justify-between items-center pb-4 pt-2\"><div class=\"flex items-center gap-2\"><button onclick=\"toggleDarkMode()\" class=\"p-2 rounded-md border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 w-10 inline-flex items-center justify-center\" aria-label=\"Toggle dark mode\"><svg id=\"sun-icon\" xmlns=\"http://www.w3.org/2000/svg\" width=\"20\" height=\"20\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"w-5 h-5\"><circle cx=\"12\" cy=\"12\" r=\"5\"></circle> <line x1=\"12\" y1=\"1\" x2=\"12\" y2=\"3\"></line> <line x1=\"12\" y1=\"21\" x2=\"12\" y2=\"23\"></line> <line x1=\"4.22\" y1=\"4.22\" x2=\"5.64\" y2=\"5.64\"></line> <line x1=\"18.36\" y1=\"18.36\" x2=\"19.78\" y2=\"19.78\"></line> <line x1=\"1\" y1=\"12\" x2=\"3\" y2=\"12\"></line> <line x1=\"21\" y1=\"12\" x2=\"23\" y2=\"12\"></line> <line x1=\"4.22\" y1=\"19.78\" x2=\"5.64\" y2=\"18.36\"></line> <line x1=\"18.36\" y1=\"5.64\" x2=\"19.78\" y2=\"4.22\"></line></svg> <svg id=\"moon-icon\" xmlns=\"http://www.w3.org/2000/svg\" width=\"20\" height=\"20\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"hidden w-5 h-5\"><path d=\"M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z\"></path></svg></button> <a href=\"/create\" class=\"p-2 rounded-md border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 w-10 inline-flex items-center justify-center\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"20\" height=\"20\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"w-5 h-5\"><line x1=\"12\" y1=\"5\" x2=\"12\" y2=\"19\"></line> <line x1=\"5\" y1=\"12\" x2=\"19\" y2=\"12\"></line></svg></a> <a href=\"https://insigh.to/b/marked\" class=\"p-2 rounded-md border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 gap-2 inline-flex items-center justify-center\"><svg viewBox=\"0 0 24 24\" fill=\"none\" width=\"20\" height=\"20\" xmlns=\"http://www.w3.org/2000/svg\"><g id=\"SVGRepo_bgCarrier\" stroke-width=\"0\"></g><g id=\"SVGRepo_tracerCarrier\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></g><g id=\"SVGRepo_iconCarrier\"><path d=\"M12 3L13.4302 8.31181C13.6047 8.96 13.692 9.28409 13.8642 9.54905C14.0166 9.78349 14.2165 9.98336 14.451 10.1358C14.7159 10.308 15.04 10.3953 15.6882 10.5698L21 12L15.6882 13.4302C15.04 13.6047 14.7159 13.692 14.451 13.8642C14.2165 14.0166 14.0166 14.2165 13.8642 14.451C13.692 14.7159 13.6047 15.04 13.4302 15.6882L12 21L10.5698 15.6882C10.3953 15.04 10.308 14.7159 10.1358 14.451C9.98336 14.2165 9.78349 14.0166 9.54905 13.8642C9.28409 13.692 8.96 13.6047 8.31181 13.4302L3 12L8.31181 10.5698C8.96 10.3953 9.28409 10.308 9.54905 10.1358C9.78349 9.98336 9.98336 9.78349 10.1358 9.54905C10.308 9.28409 10.3953 8.96 10.5698 8.31181L12 3Z\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path></g></svg> Request a feature </a></div><div class=\"flex items-center\"><input type=\"text\" id=\"shareUrl\" readonly class=\"bg-background border border-input text-foreground rounded-l-md px-3 py-2 focus:outline-none\"> <button onclick=\"copyUrl()\" class=\"bg-foreground hover:bg-primary/90 text-background px-3 py-2 rounded-r-md\" id=\"copyButton\">Copy</button></div></div><div class=\"flex justify-center\"><div class=\"w-full\"><form id=\"documentForm\"><div class=\"mb-4\"><input type=\"text\" id=\"title\" name=\"title\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(document.Title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/pages/document.templ`, Line: 62, Col: 92}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"w-full p-2 text-3xl font-semibold focus:ring-background focus:border-background bg-background outline-none appearance-none text-foreground\"></div><div class=\"mb-4\"><textarea id=\"content\" name=\"content\" class=\"bg-background\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(document.Content)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/pages/document.templ`, Line: 66, Col: 105}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</textarea></div></form><div id=\"saveIndicator\" class=\"text-sm text-foreground/80 mt-2\"></div></div></div></div><script src=\"https://cdn.jsdelivr.net/npm/easymde/dist/easymde.min.js\"></script> <style>\n            /* Override EasyMDE styles to match parent background */\n            .EasyMDEContainer {\n                background-color: inherit;\n            }\n            .EasyMDEContainer .CodeMirror {\n                background-color: inherit;\n                color: inherit;\n                border: none !important;\n\n            }\n            .EasyMDEContainer .CodeMirror-cursor {\n                border-left: 1px solid hsl(var(--foreground));\n            }\n            .EasyMDEContainer .CodeMirror-selected {\n                background: hsl(var(--secondary));\n            }\n            .EasyMDEContainer .CodeMirror-line {\n                color: inherit;\n            }\n            .EasyMDEContainer .editor-toolbar {\n                display: none;\n            }\n            .EasyMDEContainer .editor-preview {\n                background-color: inherit;\n                color: inherit;\n            }\n            .EasyMDEContainer .CodeMirror-scroll {\n                min-height: 80vh;\n            }\n        </style> <link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/easymde/dist/easymde.min.css\"><script>\n    function toggleDarkMode() {\n        const html = document.documentElement;\n        const sunIcon = document.getElementById('sun-icon');\n        const moonIcon = document.getElementById('moon-icon');\n\n        html.classList.toggle('dark');\n        const isDark = html.classList.contains('dark');\n        localStorage.setItem('darkMode', isDark);\n\n        sunIcon.classList.toggle('hidden', isDark);\n        moonIcon.classList.toggle('hidden', !isDark);\n    }\n\n    // Check for saved dark mode preference\n    if (localStorage.getItem('darkMode') === 'true') {\n        document.documentElement.classList.add('dark');\n        document.addEventListener('DOMContentLoaded', () => {\n            document.getElementById('sun-icon').classList.add('hidden');\n            document.getElementById('moon-icon').classList.remove('hidden');\n        });\n    }\n\n    const shareUrlInput = document.getElementById('shareUrl');\n    shareUrlInput.value = window.location.href;\n\n    function copyUrl() {\n        const urlInput = document.getElementById('shareUrl');\n        const copyButton = document.getElementById('copyButton');\n        urlInput.select();\n        document.execCommand('copy');\n        copyButton.textContent = 'Copied!';\n        setTimeout(() => {\n            copyButton.textContent = 'Copy';\n        }, 2000);\n    }\n\n    (function() {\n        const titleInput = document.getElementById(\"title\");\n        const contentInput = document.getElementById(\"content\");\n        const saveIndicator = document.getElementById(\"saveIndicator\");\n\n        let saveTimeout;\n        let lastSavedContent = { title: titleInput.value, content: contentInput.value };\n        let ws;\n        let isLocalUpdate = false;\n\n        const easymde = new EasyMDE({\n            element: contentInput,\n            spellChecker: false,\n            autosave: {\n                enabled: false\n            },\n            toolbar: false,\n            status: false,\n            placeholder: \"Start typing here...\",\n        });\n\n        function updateSaveIndicator(status) {\n            saveIndicator.textContent = status;\n        }\n\n        function setupWebSocket() {\n            const path = window.location.pathname;\n            const documentID = path.split('/').pop();\n            const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';\n            ws = new WebSocket(`${protocol}//${window.location.host}/ws/${documentID}`);\n\n            ws.onopen = function() {\n                console.log('WebSocket connection established');\n                updateSaveIndicator(\"Connected\");\n            };\n\n            ws.onclose = function(e) {\n                console.log('WebSocket connection closed:', e.reason);\n                updateSaveIndicator(\"Disconnected\");\n                setTimeout(setupWebSocket, 1000); // Attempt to reconnect after 1 second\n            };\n\n            ws.onerror = function(err) {\n                console.error('WebSocket error:', err);\n                updateSaveIndicator(\"Error\");\n                ws.close();\n            };\n\n            ws.onmessage = function(event) {\n                const data = JSON.parse(event.data);\n                if (data.type === \"update\" && !isLocalUpdate) {\n                    titleInput.value = data.title;\n                    if (data.content !== easymde.value()) {\n                        easymde.value(data.content);\n                    }\n                } else if (data.type === \"saved\") {\n                    lastSavedContent = { title: data.title, content: data.content };\n                    updateSaveIndicator(\"Saved\");\n                }\n            };\n        }\n\n        setupWebSocket();\n\n        function debounce(func, wait) {\n            return function(...args) {\n                clearTimeout(saveTimeout);\n                saveTimeout = setTimeout(() => func.apply(this, args), wait);\n            };\n        }\n\n        function sendUpdate(isSave = false) {\n            const currentContent = { title: titleInput.value, content: easymde.value() };\n            if (ws.readyState === WebSocket.OPEN) {\n                isLocalUpdate = true;\n                ws.send(JSON.stringify({ \n                    type: isSave ? \"save\" : \"update\", \n                    ...currentContent \n                }));\n                setTimeout(() => { isLocalUpdate = false; }, 10000); // Reset after a short delay\n                if (!isSave) {\n                    updateSaveIndicator(\"Editing...\");\n                }\n            } else {\n                console.log('WebSocket is not open. Attempting to reconnect...');\n                updateSaveIndicator(\"Reconnecting...\");\n                setupWebSocket();\n            }\n        }\n\n        const debouncedSave = debounce(() => {\n            sendUpdate(true);\n            updateSaveIndicator(\"Saving...\");\n        }, 2000); // 2 seconds delay for saving\n\n        titleInput.addEventListener(\"input\", () => {\n            sendUpdate();\n            debouncedSave();\n        });\n\n        easymde.codemirror.on(\"change\", () => {\n            sendUpdate();\n            debouncedSave();\n        });\n\n        // Initial save indicator\n        updateSaveIndicator(\"Connected\");\n\n        // Check for saved dark mode preference\n        if (localStorage.getItem('darkMode') === 'true') {\n            document.documentElement.classList.add('dark');\n            document.getElementById('sun-icon').classList.add('hidden');\n            document.getElementById('moon-icon').classList.remove('hidden');\n        }\n    })();\n</script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
