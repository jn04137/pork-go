<script>
  import { onMount, onDestroy } from 'svelte'
  import { Editor } from '@tiptap/core'
  import StarterKit from '@tiptap/starter-kit'
  import Placeholder from '@tiptap/extension-placeholder'
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
          placeholder: "Write a comment...",
          emptyEditorClass: "is-editor-empty"
        }),
        Paragraph.configure({
          
        }) 
      ],
      //content: '<p>Write your comment...</p>',
      onTransaction: () => {
        // force re-render so `editor.isActive` works as expected
        //editor = editor
        editorHTML = editor.getHTML()
      },
      editorProps: {
        attributes: {
          class: "border text-sm border-grey-300 px-2 py-1 rounded min-h-[150px] shadow-sm"
        }
      }
    })
  })

  onDestroy(() => {
    if (editor) {
      editor.destroy()
    }
  })

  const handleCreateComment = async () => {
  }

</script>

<div>
  <div>
  {#if editor}
    <button
      on:click={() => editor.chain().focus().toggleHeading({ level: 1}).run()}
      class:active={editor.isActive('heading', { level: 1 })}
      >
      H1
    </button>
    <button
      on:click={() => editor.chain().focus().toggleHeading({ level: 2 }).run()}
      class:active={editor.isActive('heading', { level: 2 })}
      >
      H2
    </button>
    <button on:click={() => editor.chain().focus().setParagraph().run()} class:active={editor.isActive('paragraph')}>
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

  :global(.tiptap p.is-editor-empty:first-child::before) {
    color: #adb5bd;
    content: attr(data-placeholder);
    float: left;
    height: 0;
    pointer-events: none;
  }

</style>
