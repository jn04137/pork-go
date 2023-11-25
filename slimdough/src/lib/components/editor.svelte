<script>
  import { onMount, onDestroy } from 'svelte'
  import { Editor } from '@tiptap/core'
  import StarterKit from '@tiptap/starter-kit'
  import Placeholder from '@tiptap/extension-placeholder'
  import Heading from '@tiptap/extension-heading'
  import Paragraph from '@tiptap/extension-paragraph'

  let element
  let editor
  export let editorHTML

  onMount(() => {
    editor = new Editor({
      element: element,
      extensions: [
        StarterKit,
        Placeholder.configure({
          placeholder: "Write a post..."
        }),
        Heading.configure({
          HTMLAttributes: {
            class: 'text-[18px] pt-2 pb-1'
          }
        }),
        Paragraph.configure({
          HTMLAttributes: {
            class: 'text-[14px] font-light'
          }
        })
      ],
      onTransaction: () => {
        // force re-render so `editor.isActive` works as expected
        editor = editor
        editorHTML = editor.getHTML()
      },
      editorProps: {
        attributes: {
          class: "border border-grey-300 px-2 py-1.5 rounded-b-lg min-h-[150px]"
        }
      }
    })
  })

  onDestroy(() => {
    if (editor) {
     editor.destroy()
    }
  })

  let editorButtons = "w-8 py-2"
  let firstEditorButton = "rounded-tl-lg w-8 py-2"

</script>

<div>
  <div class="border rounded-t-lg border-b-0">
  {#if editor}
    <button
      on:click={() => editor.chain().focus().toggleHeading({ level: 1}).run()}
      class:active={editor.isActive('heading', { level: 1 })}
      class={firstEditorButton}
      >
      H1
    </button>
    <button
      on:click={() => editor.chain().focus().toggleHeading({ level: 2 }).run()}
      class:active={editor.isActive('heading', { level: 2 })}
      class={editorButtons}
      >
      H2
    </button>
    <button 
      on:click={() => editor.chain().focus().setParagraph().run()} 
      class:active={editor.isActive('paragraph')}
      class={editorButtons}
    >
      P
    </button>
  {/if}
  </div>

  <div bind:this={element} />
</div>

<style>
 button.active {
   background: black;
   color: white;
 }
</style>
