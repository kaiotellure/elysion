// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/ikaio/tailmplx/models"
import "github.com/ikaio/tailmplx/components"

func FilestoreUpload() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col-reverse sm:flex-row gap-4 w-full\"><form class=\"flex flex-col gap-2\" id=\"file-upload-form\" hx-encoding=\"multipart/form-data\" hx-post=\"#\" hx-swap=\"outerHTML\"><label for=\"title\">Upload Name</label> <input class=\"outline-none rounded border-2 px-4 py-2 border-violet-800\" name=\"title\" type=\"text\" placeholder=\"Describe what is this upload for.\" maxlength=\"256\"> <label for=\"author\">Author Identifier</label> <input class=\"outline-none rounded border-2 px-4 py-2 border-violet-800\" name=\"author\" type=\"text\" placeholder=\"Your name or social tag\" maxlength=\"24\"> <input class=\"border-violet-800 border-4 border-dashed px-20 py-10\" type=\"file\" multiple name=\"attachments\"> <button id=\"publish\" class=\"bg-violet-800 hover:bg-violet-700 px-3 py-2 w-fit rounded font-semibold text-white\">Publish Upload</button></form><div class=\"p-3 rounded bg-black/5\"><span class=\"font-bold text-current/75\">NOTE:</span> This is not a permanent file archive solution, your files can be deleted at any time.<br><span class=\"font-bold text-current/75\">NOTE:</span> The max file size allowed is 100MB, upload will abort if longer.<br><span class=\"font-semibold text-sm text-black/50\">By uploading you are agreeing with our community guidelines.</span><br><br><span class=\"font-bold text-current/75\">1.</span> I'm not uploading copyright holded content.<br><span class=\"font-bold text-current/75\">2.</span> I'm not uploading content promoting violence against minors or animals.<br><span class=\"font-bold text-current/75\">3.</span> Explicit content is allowed if under moderated quantity and should not involve minors or animals.<br></div></div><script>\n        htmx.on('#file-upload-form', 'htmx:xhr:progress', function(e) {\n\t\t\tlet percent = Math.floor(e.detail.loaded/e.detail.total * 100);\n\t\t\tlet publish_button = htmx.find('#publish');\n          \tpublish_button.textContent = `Uploading ${percent}%`;\n\t\t\tpublish_button.style.cursor = \"not-allowed\";\n\t\t\tpublish_button.disabled = true;\n        });\n    </script><span class=\"mt-4 font-bold\">Recent Uploads</span><div class=\"flex flex-col gap-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if uploads, err := models.ListUploads(100); err == nil && len(uploads) != 0 {
			for i := len(uploads) - 1; i >= 0; i-- {
				templ_7745c5c3_Err = components.DisplayUpload(uploads[i]).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
