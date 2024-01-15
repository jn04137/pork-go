import Layout from '../../components/layout'
import { useNavigate } from 'react-router-dom'
import axios from 'axios'
import { EditorContent, useEditor } from '@tiptap/react'
import StarterKit from '@tiptap/starter-kit'
import Placeholder from '@tiptap/extension-placeholder'
import React, { useState, Dispatch, SetStateAction } from 'react'

import "./custom.css"

function CreatePostPage() {
  const navigate = useNavigate()

  const [body, setBody] = useState("")
  const [title, setTitle] = useState("")

  let postData = {
    title: title,
    body: body
  }

  const handlePublish = async (e: React.MouseEvent) => {
    e.preventDefault()
    if(postData.title.length === 0) {
      alert("Please enter a title")
      throw new Error("Title wasn't supplied")
    };
    try {
      await axios.post(`${import.meta.env.VITE_API_URL}/api/protected/createpost`, postData, {
        withCredentials: true
      })
    } catch(e) {
      console.error(e)
    }
    navigate("/", { replace: true })
  }

  return(
    <Layout>
      <div className="flex justify-center pt-10">
        <form className="bg-white border rounded p-4 w-[650px] space-y-1 shadow">
          <input 
            placeholder="title" 
            className="text-xl rounded-t-lg py-0.5 px-2 border border-grey-300 w-full"
            onChange={e => setTitle(e.target.value)}
            required
          />
          <TipTap setText={setBody}/>
          <div className="flex justify-end">
            <button
              onClick={(e) => handlePublish(e)}
              className="bg-blue-500 text-white text-sm px-4 py-1 rounded shadow">
              Publish
            </button>
          </div>
        </form>
      </div>

    </Layout>
  )
}

function TipTap({setText}: {setText: Dispatch<SetStateAction<string>>}) {
  const extensions = [
    StarterKit,
    Placeholder.configure({
      placeholder: "Start your post..."
    })
  ]
  
  const editor = useEditor({
      extensions,
      editorProps: {
        attributes: {
          class: 'border border-grey-300 px-2 py-1.5 rounded-b-lg min-h-[150px]'
        }
      },
      onUpdate: () => {setText(editor!.getHTML())},
  })

  return <EditorContent editor={editor} />

}

export default CreatePostPage
