import {
  useLoaderData,
  Params,
} from 'react-router-dom'
import axios from 'axios'
import Layout from '../components/layout'
import { IPostData } from './home'
import {
} from '@tiptap/react'
import Placeholder from '@tiptap/extension-placeholder'
import React, {
  useState,
} from 'react'
import { useInfiniteQuery, useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { 
  EditorContent,
  useEditor
} from '@tiptap/react'
import StarterKit from '@tiptap/starter-kit'
import { HiChevronDown, HiChevronUp } from 'react-icons/hi'

import { loggedInQuery } from '../components/navbar'
import DropDownPostMenu from '../components/DropDownPostMenu'
import DropDownCommentMenu from '../components/DropDownCommentMenu'

const apiURL = import.meta.env.VITE_API_URL
const cardBgColor = "[#2e2e2e]"

interface IPostContent {
  content: IPostData;
}

interface ICommentData {
	CommentId: number;
  Owner: string;
  Body: string;
  CreatedAt: string;
}

interface ICreateComment {
  body: string
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
  const queryClient = useQueryClient()

  const commentData = {
    body: newComment
  }

  const extensions = [
    StarterKit,
    Placeholder.configure({
      placeholder: "Write your comment..."
    })
  ]

  const commentEditor = useEditor({
    extensions,
    editorProps: {
    attributes: {
        class: 'py-1 rounded min-h-[60px]'
    }
    },
    onUpdate: () => {setNewComment(commentEditor!.getHTML())},
  })

  const createComment = async ({commentData}: {
    commentData: ICreateComment
  }) => {
    try{
      const response = await axios.post(`${import.meta.env.VITE_API_URL}/api/protected/createcomment/${content.PostId}`, commentData, {
        withCredentials: true
      })
      return response
    } catch(e) {
      console.error(e)
    }
  }

  const createCommentMutation = useMutation({
    mutationFn: createComment,
    onSuccess: () => {
      setTimeout(() => {
        queryClient.invalidateQueries({ queryKey: ['commentFeed', content.PostId]})
      }, 2000)
      commentEditor?.commands.clearContent();
    }
  })

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
            <CommentEditor editor={commentEditor}/>
            <div className="flex justify-end pt-2">
              <button
                className="dark:bg-blue-600 text-white text-sm px-2 py-1 rounded"
                onClick={() => createCommentMutation.mutate({commentData})}
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

function CommentEditor({editor}: {
  editor: any
}) {
  return (
    <>
      <h1 className="text-xs">Share your comment</h1>
      <EditorContent editor={editor} />
    </>
  )

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
									commentId={comment.CommentId}
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
	commentId: number;
  body: string;
  author: string;
  date: string;
}

function CommentCard({commentId, body, author, date}: ICommentCardProps) {

	const { data, isSuccess } = loggedInQuery()
	const [editable, setEditable] = useState(false)

  return(
    <div className="dark:bg-[#2e2e2e] rounded px-5 py-2 shadow">
      <div className='text-sm py-2'>
        <TipTap content={body} editable={editable}/>
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
				{
					isSuccess
					&& data.username === author
					&& <DropDownCommentMenu 
						commentId={commentId} 
						editable={editable} 
						setEditable={setEditable} />
				}
			</div>
    </div>
  )
}

function TipTap({content, editable}: {
	content: string,
	editable?: boolean
}) {
  const extensions = [
    StarterKit,
  ]
	const editor = useEditor({
		extensions,
		content,
	})

	if(editable !== undefined) {
		editor?.setEditable(editable)
	}

  return (
		<EditorContent editor={editor}/>
  )
}

interface IPostCardData {
  title: string;
  body: string;
  author: string;
  date: string;
  postId: number;
}

export function PostCard({title, body, author, date, postId}: IPostCardData) {
  //let instance = new Date()
  
  const { data, isSuccess } = loggedInQuery()
  
  return(
    <div className={`dark:bg-${cardBgColor} rounded px-4 py-2`}>
      <div className='flex justify-between items-center text-[24px] font-semibold'>
        <div>
          {title}
        </div>
        {
          isSuccess 
          && data.username === author
          && <DropDownPostMenu postId={postId}/>
        }
      </div>
      <div className='pt-1 pb-2 text-sm h-fit'>
				<TipTap content={body} />
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
        <PostCardPoints postId={postId}/>
      </div>
    </div>
  )
}

function PostCardPoints({postId}: {postId: number}) {
  type IUserPointing = {
    pointing: string
  }
  
  const queryClient = useQueryClient()

  const getPoints = async () => {
    try {
      const response = await axios.get(`${apiURL}/api/public/post/points/${postId}`, {
        withCredentials: true
      })
      return response.data
    } catch (e) {
      console.log(e)
    }
  }

  const mutatePointsAPI = async (userPointing: IUserPointing) => {
    try {
      const response = await axios.post(`${apiURL}/api/protected/mutatepostpoint/${postId}`,
      userPointing, {
        withCredentials: true
      })
      return response.data
    } catch (e) {
      console.log(e)
    }
  }

  const pointsQuery = useQuery({
    queryKey: ['postPoints', postId],
    queryFn: getPoints
  })

  const userPointingMutation = useMutation({
    mutationFn: mutatePointsAPI,
    onSuccess: (data) => {
      console.log(data)
      queryClient.setQueryData(['postPoints', postId], data)
    }
  })

  return(
    <div className="flex justify-between divide-x items-center border rounded">
      <button 
        className="flex items-center px-2 h-full hover:bg-blue-600"
        onClick={() => userPointingMutation.mutate({pointing: 'plus'})}
      >
        <HiChevronUp size={20}/>
      </button>
        <div className="flex items-center w-16 items-center text-center h-full">
          <div className="text-center w-full">
            {pointsQuery.isSuccess ? pointsQuery.data.points : 0}
          </div>
        </div>
      <button 
        className="flex items-center px-2 h-full hover:bg-blue-600"
        onClick={() => userPointingMutation.mutate({pointing: 'minus'})}
      >
        <HiChevronDown size={20}/>
      </button>
    </div>
  )
}

export default ViewPostPage
