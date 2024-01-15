import Layout from '../components/layout'
import React from 'react'
import { Link } from 'react-router-dom'
import { useInfiniteQuery } from '@tanstack/react-query'
import axios from 'axios'
import { EditorProvider, FloatingMenu, BubbleMenu } from '@tiptap/react'
import StarterKit from '@tiptap/starter-kit'

const endpoint = "/api/public/getfeedposts/"
const fetchFeed = async ({ pageParam }: {pageParam: number}) => {
  const response = await axios.get(endpoint+pageParam)
  return response.data
}

function Home() {
  return(
    <Layout>
      <div className="w-full">
        <div className="flex space-x-2">
          <Feed/>
          <div className="bg-white rounded grow h-fit p-2">
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
   //isFetching,
   isFetchingNextPage,
   status
  } = useInfiniteQuery({
    queryKey: ['userFeed'],
    queryFn: fetchFeed,
    initialPageParam: 0,
    getNextPageParam: (lastPage) => lastPage.nextCursor ?? undefined
  })

  return(
    <div className="space-y-2.5 w-[650px]">
      <div className="px-2 py-1 rounded bg-white text-xl border shadow">Feed</div>
      <div className="border shadow bg-white rounded px-2 py-2">
        <Link to="/createpost">
          <div
            className="border border-grey-300 px-2 py-1 text-[#C0C0C0] bg-[#F8F8F8] rounded-lg border-2 
            hover:cursor-text hover:bg-white">
            Create a post
          </div>
        </Link>
      </div>
      {status === 'pending' ? (
        <p>Loading...</p>
      ) : status === 'error' ? (
        <p>Error: {error.message}</p>
      ) : (
        <>
          {data.pages.map((pages, i) => (
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
  postId?: number;
}

export function PostCard({title, body, author, date, postId}: IPostCardData) {
  //let instance = new Date()
  
  return(
    <div className="bg-white border rounded px-4 py-2 shadow">
      <div className='text-[24px] font-semibold'>
      <Link to={`/viewpost/${postId}`}>
        {title}
      </Link>
      </div>
      <div className='pt-1 pb-2 text-sm h-fit'>
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

export default Home
