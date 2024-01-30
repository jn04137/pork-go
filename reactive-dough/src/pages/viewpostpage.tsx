import { useLoaderData, Params } from 'react-router-dom'
import axios from 'axios'
import Layout from '../components/layout'
import { PostCard, IPostData } from './home'
import { EditorContent, useEditor } from '@tiptap/react'
import Placeholder from '@tiptap/extension-placeholder'
import React, { useState, Dispatch, SetStateAction } from 'react'
import { useInfiniteQuery } from '@tanstack/react-query'
import { EditorProvider, FloatingMenu, BubbleMenu } from '@tiptap/react'
import StarterKit from '@tiptap/starter-kit'

interface IPostContent {
  content: IPostData;
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

function ViewPostPage() {
  const { content } = useLoaderData() as IPostContent
  const [newComment, setNewComment] = useState<string>("")

  const commentData = {
    body: newComment
  }

  const handleCreateComment = async (e: React.MouseEvent) => {
    e.preventDefault()
    try{
      const response = await axios.post(`${import.meta.env.VITE_API_URL}/api/protected/createcomment/${content.PostId}`, commentData, {
        withCredentials: true
      })
      return response
    } catch(e) {
      console.error(e)
    }
  } 

  return (
    <Layout>
      <div className="flex justify-center w-full">
        <div className="w-[650px] space-y-2">
          <PostCard
            title={content.Title}
            body={content.Body}
            date={content.CreatedAt}
            author={content.Owner}
            postId={content.PostId}
          />
          <div className="dark:bg-[#2e2e2e] px-5 py-3">
            <CommentEditor setText={setNewComment}/>
            <div className="flex justify-end pt-2">
              <button
                className="dark:bg-blue-600 text-white text-sm px-2 py-1 rounded"
                onClick={(e) => handleCreateComment(e)}
              >
                Create Comment
              </button>
            </div>
          </div>
          <CommentFeed
            postId={content.PostId}
          />
        </div>
      </div>
    </Layout>
  )
}

function CommentEditor({setText}: {setText: Dispatch<SetStateAction<string>>}) {
  const extensions = [
    StarterKit,
    Placeholder.configure({
      placeholder: "Write your comment..."
    })
  ]

  const editor = useEditor({
      extensions,
      editorProps: {
        attributes: {
          class: 'py-1 rounded min-h-[60px]'
        }
      },
      onUpdate: () => {setText(editor!.getHTML())},
  })

  return (
    <>
      <h1 className="text-xs">Share your comment</h1>
      <EditorContent editor={editor} />
    </>
  )

}

interface ICommentData {
  Owner: string;
  Body: string;
  CreatedAt: string;
}

function CommentFeed({postId}: {postId: number}) {

  const fetchComments = async({ pageParam=0 }) => {
    try {
      const response = await axios.get(`${import.meta.env.VITE_API_URL}/api/public/loadcomments/${postId}/${pageParam}`)
      return response.data
    } catch(err) {
      console.error(err)
    }
  }

  const {
   data,
   error,
   fetchNextPage,
   hasNextPage,
   //isFetching,
   isFetchingNextPage,
   status
  } = useInfiniteQuery({
    queryKey: ['commentFeed', postId],
    queryFn: fetchComments,
    initialPageParam: 0,
    getNextPageParam: (lastPage) => lastPage.nextCursor ?? undefined
  })


  return(
    <div className="space-y-2">
      {status === 'pending' ? (
        <p>Loading...</p>
      ) : status === 'error' ? (
        <p>Error: {error.message}</p>
      ): (
        <>
          {data.pages.map((pages, i) => (
            <React.Fragment key={i}>
              {pages.comments.map((comment: ICommentData, j: number) => (
                <CommentCard
                  author={comment.Owner}
                  body={comment.Body}
                  date={comment.CreatedAt}
                  key={j}
                />
              ))}
            </React.Fragment>
          ))}
        </>
      )}
      <div className="w-full flex justify-center">
        <button
          onClick={() => fetchNextPage()}
          disabled={!hasNextPage || isFetchingNextPage}
          className="text-sm bg-blue-500 text-white rounded py-0.5 px-4"
        >
          {isFetchingNextPage
            ? 'Loading more...'
            : hasNextPage
              ? 'Load More'
              : 'Nothing more to load'}
        </button>
      </div>
    </div>
  )
}

interface ICommentCardProps {
  body: string;
  author: string;
  date: string;
}

function CommentCard({body, author, date}: ICommentCardProps) {
  return(
    <div className="dark:bg-[#2e2e2e] rounded px-5 py-4 shadow">
      <div className='text-sm'>
        <TipTap content={body}/>
      </div>
      <div className='flex flex-col text-xs'>
        <div>
          {author}
        </div>
        <div className='italic'>
          {new Date(date).toLocaleString()}
        </div>
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
      extensions={extensions}
      content={content}
      editable={false}
    >
      <BubbleMenu>This is the bubble menu</BubbleMenu>
      <FloatingMenu>This is the floating menu</FloatingMenu>
    </EditorProvider>
  )
}

export default ViewPostPage
