<script>
  import Layout from '../lib/components/layout.svelte'
  import PostCard from '../lib/components/post_card.svelte'
  import CommentFeed from '../lib/components/comment_feed.svelte'
  import CommentEditor from '../lib/components/comment_editor.svelte'

  import axios from 'axios'
  import { onMount } from 'svelte'

  export let postId
  let post

  let createCommentData = {
    body: "",
  }

  onMount(async () => {
    try {
      const response = await axios.get(`/public/viewpost/${postId}`)
      post = response.data.post
    } catch(err) {
      console.error(err)
    }
  })

  const handleCreatePost = async () => {
    try{
      const response = await axios.post(`/api/createcomment/${postId}`, createCommentData, {
        withCredentials: true
      })
    } catch(e) {
      console.error(e)
    }
  } 

</script>

<Layout>
  <body slot="content" class="w-full">
    {#if post === undefined}
      <PostCard/>
    {:else}
      <div class="flex justify-center">
        <div class="w-[650px] space-y-2">
          <PostCard>
            <h1 slot="title">{post.Title}</h1>
            <p slot="body">{@html JSON.parse(post.Body)}</p>
            <div slot="author">{post.Owner}</div>
            <div slot="date">{new Date(post.CreatedAt).toLocaleString()}</div>
          </PostCard>
          <div class="bg-white p-4 rounded shadow-lg space-y-2">
            <CommentEditor bind:editorHTML={createCommentData.body}/>
            <div class="flex justify-end">
              <button 
                class="bg-blue-500 text-white text-sm px-2 py-0.5 rounded-2xl shadow"
                on:click|preventDefault={() => handleCreatePost()}
              >
                Create Comment
              </button>
            </div>
          </div>
          <CommentFeed postId={postId}/>
        </div>
      </div>
    {/if}
  </body>
</Layout>

<style></style>
