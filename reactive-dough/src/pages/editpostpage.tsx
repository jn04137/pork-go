import { useEffect } from 'react'
import Layout from '../components/layout'
import axios from 'axios'
import { 
	useNavigate, 
	Params, 
	useLoaderData 
} from 'react-router-dom'

import { EditorContent, useEditor } from '@tiptap/react'
import StarterKit from '@tiptap/starter-kit'

import { 
	useState, 
	Dispatch, 
	SetStateAction 
} from 'react'

import { loggedInQuery } from '../components/navbar'

interface IPostContent {
  content: IPostData;
}

export interface IPostData {
  PostId: number;
  Title: string;
  Owner: string;
  Body: string;
  CreatedAt: string
}

export async function loader({params}: {params: Params<"postId">}) {
  try {
    const post = await axios.get(`${import.meta.env.VITE_API_URL}/api/public/viewpost/${params.postId}`)
    const content = post.data.post
    return { content }
  } catch(error) {
    console.error(error)
  }
}

function EditPostPage() {
	const { content } = useLoaderData() as IPostContent
  const [body, setBody] = useState(content.Body)
  const [title, setTitle] = useState(content.Title)
  const navigate = useNavigate()
	const { data, isSuccess } = loggedInQuery()

	useEffect(() => {
		if(isSuccess && data.username !== content.Owner) {
			navigate("/", { replace: true })
		}
	})
  
  let postData = {
    title: title,
    body: body,
		postId: content.PostId
  }

	const handlePublish = async (e: React.MouseEvent) => {
    e.preventDefault()
    if(postData.title.length === 0) {
      alert("Please enter a title")
      throw new Error("Title wasn't supplied")
    };
    try {
      await axios.post(`${import.meta.env.VITE_API_URL}/api/protected/post/edit`, postData, {
        withCredentials: true
      })
    } catch(e) {
      console.error(e)
    }
    navigate(`/viewpost/${content.PostId}`, { replace: true })
  }

	return(
		<Layout>
			<form className="dark:bg-[#2e2e2e] rounded p-4 w-[650px] space-y-1">
				<input 
					placeholder="enter title" 
					value={title}
					className="dark:bg-[#2e2e2e] text-xl py-0.5 px-2 w-full focus:outline-none border-b"
					onChange={e => setTitle(e.target.value)}
					required
				/>
				<TipTap setText={setBody} text={body}/>
				<div className="flex justify-end">
					<button
						onClick={(e) => handlePublish(e)}
						className="bg-blue-500 text-white text-sm px-4 py-1 rounded shadow">
						Publish
					</button>
				</div>
			</form>
		</Layout>
	)
}

function TipTap({setText, text}: {
	setText: Dispatch<SetStateAction<string>>, 
	text: string
}) {

  const extensions = [
    StarterKit,
  ]
  
  const editor = useEditor({
		extensions,
		content: text,
		editorProps: {
			attributes: {
				class: 'px-2 py-1.5 rounded-b-lg min-h-[150px]'
			}
		},
		onUpdate: () => {setText(editor!.getHTML())},
  })

  return <EditorContent editor={editor} />

}

export default EditPostPage

