import Layout from '../components/layout'
import { Link } from 'react-router-dom'
import { EditorProvider } from '@tiptap/react';
import StarterKit from '@tiptap/starter-kit';

function MyProfilePage() {

  return(
    <Layout>
      <div>
        <h1 className="text-2xl">My Profile</h1>
      </div>
				<p className="h-full grow items-center flex items-center">
					Profile Page coming soon 
				</p>
    </Layout>
  )
}

interface IPostCardData {
  title: string;
  body: string;
  author: string;
  date: string;
  postId: number;
}


const cardBgColor = "[#2e2e2e]"
export function PostCard({title, body, author, date, postId}: IPostCardData) {
  
  return(
    <div className={`dark:bg-${cardBgColor} rounded px-4 py-2`}>
      <div className='text-[24px] font-semibold'>
      <Link to={`/viewpost/${postId}`}>
        {title}
      </Link>
      </div>
      <div className='pt-1 pb-2 text-sm h-fit'>
        <TipTap content={body}/>
      </div>
      <div className="flex justify-between">
        <div className='flex flex-col text-xs'>
          <div>
            {author}
          </div>
          <div className='italic'>
            {new Date(date).toLocaleString()}
          </div>
        </div>
        {/*
          <PostCardPoints postId={postId}/>
        */}
      </div>
    </div>
  )
}

function TipTap({content}: {content: string}) {
  const extensions = [
    StarterKit
  ]

  return (
    <EditorProvider 
      children={""}
      extensions={extensions}
      content={content}
      editable={false}
    />
  )
}

export default MyProfilePage
