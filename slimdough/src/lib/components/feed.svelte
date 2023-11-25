<script>
  import PostCard from './post_card.svelte'
  import { Link } from 'svelte-routing'
  import { createInfiniteQuery } from '@tanstack/svelte-query'
  import axios from 'axios'

  const endpoint = "/public/getfeedposts/"
  const fetchFeed = async ({ pageParam=0 }) => {
    const response = await axios.get(endpoint+pageParam)
    return response.data
  }

  const query = createInfiniteQuery({
    queryKey: ['userFeed'],
    queryFn: ({pageParam}) => fetchFeed({pageParam}),
    getNextPageParam: (lastPage) => {
      if (lastPage.nextCursor === null) { return undefined }
      return lastPage.nextCursor
    }
  })

</script>

<div class='space-y-3 w-[650px]'>
  <div class='rounded bg-white p-3 text-xl'>Feed</div>
  <div class="bg-white rounded px-4 py-2">
    <Link to="/createpost">
      <div
        class="border border-grey-300 px-2 py-1 text-[#C0C0C0] bg-[#F8F8F8] rounded-lg border-2 
        hover:cursor-text hover:bg-white">
        Create a post
      </div>
    </Link>
  </div>
  {#if $query.isSuccess}
    {#each $query.data.pages as {posts}}
      {#each posts as post}
        <PostCard>
          <div slot="title"><Link to={`viewpost/${post.PostId}`}>{post.Title}</Link></div>
          <p slot="body"><Link to={`viewpost/${post.PostId}`}>{@html JSON.parse(post.Body)}</Link></p>
          <div slot="author">{post.Owner}</div>
          <div slot="date">{new Date(post.CreatedAt).toLocaleString()}</div>
        </PostCard>
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
        Loading...
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
