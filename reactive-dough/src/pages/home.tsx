import Layout from '../components/layout'
import React from 'react'
import { Link } from 'react-router-dom'
import { 
  useInfiniteQuery, 
  useQuery, 
  useMutation, 
  useQueryClient 
} from '@tanstack/react-query'
import axios from 'axios'
import { EditorProvider, FloatingMenu, BubbleMenu } from '@tiptap/react'
import StarterKit from '@tiptap/starter-kit'
import { HiChevronDown, HiChevronUp } from 'react-icons/hi'

const cardBgColor = "[#2e2e2e]"

const fetchFeed = async ({ pageParam }: {pageParam: number}) => {
  const response = await axios.get(`${import.meta.env.VITE_API_URL}/api/public/getfeedposts/${pageParam}`)
  return response.data
}

const apiURL = import.meta.env.VITE_API_URL

function Home() {
  return(
    <Layout>
      <div className="w-full">
        <div className="flex space-x-2 pt-3">
          <Feed/>
          <div className="rounded grow h-fit p-2">
            <h1 className="text-2xl">Welcome!</h1>
            <p className="text-sm">Make yourself at home</p>
          </div>
        </div>
      </div>
    </Layout>
  )
}

export interface IPostData {
  PostId: number;
  Title: string;
  Owner: string;
  Body: string;
  CreatedAt: string
}

function Feed() {
  const {
   data,
   error,
   fetchNextPage,
   hasNextPage,
   isFetchingNextPage,
   status
  } = useInfiniteQuery({
    queryKey: ['userFeed'],
    queryFn: fetchFeed,
    initialPageParam: 0,
    getNextPageParam: (lastPage) => lastPage.nextCursor ?? undefined
  })

  return(
    <div className="space-y-2.5 w-[550px]">
      {/* }<div className={`px-4 py-2.5 rounded text-xl dark:bg-${cardBgColor}`}>Feed</div> {*/}
      <div className={`rounded text-xl`}>Feed</div>
      <div className={`rounded px-2 py-2 dark:bg-${cardBgColor}`}>
        <Link to="/createpost">
          <div
            className="px-2 py-1 text-[#C0C0C0] rounded
            bg-[#212121] hover:cursor-text dark:hover:bg-[#121212]">
            Create a post
          </div>
        </Link>
      </div>
      {status === 'pending' ? (
        <p>Loading...</p>
      ) : status === 'error' ? (
        <p>Error: {error.message}</p>
      ) : data === null ? (
        <p>There is no data</p>
      ) : (
        <>
          { data.pages.map((pages, i) => (
            <React.Fragment key={i}>
              {pages.posts.map((post: IPostData) => (
                <PostCard
                  key={post.PostId}
                  title={post.Title}
                  author={post.Owner}
                  body={post.Body}
                  date={post.CreatedAt} 
                  postId={post.PostId}
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

interface IPostCardData {
  title: string;
  body: string;
  author: string;
  date: string;
  postId: number;
}

export function PostCard({title, body, author, date, postId}: IPostCardData) {
  //let instance = new Date()
  
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
        <PostCardPoints postId={postId}/>
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

export default Home
