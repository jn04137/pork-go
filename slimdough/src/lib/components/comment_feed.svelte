<script>
  import CommentCard from './comment_card.svelte'
  import axios from 'axios'
  import { createInfiniteQuery } from '@tanstack/svelte-query'

  export let postId

  // TODO Have to modify the endpoint to support cursor for InfiniteQuery
  const fetchComments = async({ pageParam=0 }) => {
    try {
      const response = await axios.get(`/public/loadcomments/${postId}/${pageParam}`)
      return response.data
    } catch(err) {
      console.error(err)
    }
  }
  const query = createInfiniteQuery({
    queryKey: ['postComments', 'postId'],
    queryFn: ({pageParam}) => fetchComments({pageParam}),
    getNextPageParam: (lastPage) => {
      if (lastPage.nextCursor === null) { return undefined }
      return lastPage.nextCursor
    }
  })
</script>

<div class="space-y-2">
  {#if $query.isSuccess}
    {#each $query.data.pages as {comments}}
      {#each comments as comment}
        <CommentCard>
          <p slot="body">{@html comment.Body}</p>
          <div slot="author">{comment.Owner}</div>
          <div slot="date">{new Date(comment.CreatedAt).toLocaleString()}</div>
        </CommentCard>
      {/each}
    {/each}
  {/if}
  <div class="w-full flex justify-center">
    <button
      on:click={() => $query.fetchNextPage()}
      disabled={!$query.hasNextPage || $query.isFetchingNextPage}
      class="text-sm bg-blue-500 text-white rounded-2xl py-0.5 px-4"
    >
      {#if $query.isFetching}

      {:else if $query.hasNextPage}
        Load More
      {:else}
        Nothing more to load
      {/if}
    </button>
  </div>
</div>

<style>

</style>
