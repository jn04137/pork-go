import {
  useLoaderData,
  Params,
  useNavigate
} from 'react-router-dom'
import axios from 'axios'
import Layout from '../components/layout'
import { IPostData } from './home'
import {
  EditorContent,
  useEditor
} from '@tiptap/react'
import Placeholder from '@tiptap/extension-placeholder'
import React, {
  useState,
  Fragment,
  useRef
} from 'react'
import { useInfiniteQuery, useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { EditorProvider, FloatingMenu, BubbleMenu } from '@tiptap/react'
import StarterKit from '@tiptap/starter-kit'
import { HiChevronDown, HiChevronUp } from 'react-icons/hi'
import { PiDotsThreeOutlineFill } from "react-icons/pi";
import {
  Menu,
  Dialog,
  Transition,
} from '@headlessui/react'

import { loggedInQuery } from '../components/navbar'

const apiURL = import.meta.env.VITE_API_URL
const cardBgColor = "[#2e2e2e]"

interface IPostContent {
  content: IPostData;
}

interface ICommentData {
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

function DropDownPostMenu({postId}: {
  postId: number
}) {
  const [open, setOpen] = useState(false)
	const navigate = useNavigate()

  // Send a request to the backend to delete the post
  async function handleDeleteClick(e: React.SyntheticEvent) {
    e.preventDefault()
    try {
      setOpen(true)
    } catch(e) {
      console.error(e)
    }
  }

	async function handleEditClick(e: React.SyntheticEvent) {
		e.preventDefault()
		try {
			navigate(`/editpost/${postId}`, { replace: true })
		} catch(e) {
			console.error(e)
		}
	}

  return(
    <Menu as="div" className="relative inline-block text-left text-sm">
      <div>
        <Menu.Button
          className="inline-flex w-full justify-center gap-x-1.5 rounded-md px-3 py-1
          text-sm font-semibold ring-1 ring-inset ring-gray-300 
          dark:hover:bg-[#121212]"
        >
          <PiDotsThreeOutlineFill/>
        </Menu.Button>
      </div>
      <Transition
        as={Fragment}
      >
        <Menu.Items
          className={`absolute right-0 z-10 mt-2 w-24 origin-top-right rounded-md
          ring-1 ring-black ring-opacity-5 focus:outline-none bg-[#121212] px-2 py-2 space-y-2`}
        >
          <Menu.Item>
            <a 
              className="block hover:cursor-pointer"
							onClick={handleDeleteClick}
            >Delete Post</a>
          </Menu.Item>
          <Menu.Item>
            <a 
              className="block hover:cursor-pointer"
							onClick={handleEditClick}
            >Edit Post</a>
          </Menu.Item>
        </Menu.Items> 
      </Transition>
      <ConfirmDeletePostDialog open={open} setOpen={setOpen} postId={postId}/>
    </Menu>
  )
}

function ConfirmDeletePostDialog({open, setOpen, postId}: {
    open: boolean,
    setOpen: any,
    postId: number
}) {
  const cancelButtonRef = useRef()
  const navigate = useNavigate()

	async function handleDeleteClick(e: React.MouseEvent<HTMLElement>) {
		e.preventDefault()

    try {
			console.log("This is the postId:", postId)
      const response = await axios.delete(`${apiURL}/api/protected/post/delete`, {
				withCredentials: true,
				data: {
					postId: postId
				}
      }) 
      if(response.status === 200) navigate("/", { replace: true })
    } catch(e) {
      console.log(e)
    }
  }

  return(
    <Transition.Root show={open} as={Fragment}>
      <Dialog as="div" className="relative z-10" initialFocus={cancelButtonRef} onClose={setOpen}>
        <Transition.Child
          as={Fragment}
          enter="ease-out duration-300"
          enterFrom="opacity-0"
          enterTo="opacity-100"
          leave="ease-in duration-200"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
        >
          <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
        </Transition.Child>

        <div className="fixed inset-0 z-10 w-screen overflow-y-auto">
          <div className="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
            <Transition.Child
              as={Fragment}
              enter="ease-out duration-300"
              enterFrom="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
              enterTo="opacity-100 translate-y-0 sm:scale-100"
              leave="ease-in duration-200"
              leaveFrom="opacity-100 translate-y-0 sm:scale-100"
              leaveTo="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
            >
              <Dialog.Panel className="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg">
                <div className="bg-[#212121] px-4 pb-4 pt-5 sm:p-6 sm:pb-4">
                  <div className="sm:flex sm:items-start">
                    <div className="mt-3 text-center sm:ml-4 sm:mt-0 sm:text-left">
                      <Dialog.Title as="h3" className="text-base font-semibold leading-6 text-white">
                        Delete Post
                      </Dialog.Title>
                      <div className="mt-2">
                        <p className="text-sm text-[#AAAAAA]">
                          Are you sure you want to delete this post?<br/>This will be a permanent action.
                        </p>
                      </div>
                    </div>
                  </div>
                </div>
                <div className="bg-[#3f3f3f] px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6">
                  <button
                    type="button"
                    className="inline-flex w-full justify-center rounded-md bg-red-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-red-500 sm:ml-3 sm:w-auto"
                    onClick={handleDeleteClick}
                  >
                    Delete
                  </button>
                  <button
                    type="button"
                    className="mt-3 inline-flex w-full justify-center rounded-md bg-[#5c5c5c] px-3 py-2 text-sm font-semibold text-white hover:bg-[#212121] sm:mt-0 sm:w-auto"
                    onClick={() => setOpen(false)}
                    ref={cancelButtonRef}
                  >
                    Cancel
                  </button>
                </div>
              </Dialog.Panel>
            </Transition.Child>
          </div>
        </div>
      </Dialog>
    </Transition.Root>
  )
}

export default ViewPostPage
