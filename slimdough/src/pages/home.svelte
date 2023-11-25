<script>
  import Layout from '../lib/components/layout.svelte'
  import Feed from '../lib/components/feed.svelte'
  import ProfileCard from '../lib/components/profile.svelte'
</script>

<Layout>
  <div slot="content" class='w-full'>
    <div class='flex space-x-2'>
      <Feed/>
      <!--ProfileCard/-->
      <div class="bg-white rounded grow h-fit p-2">
        <h1 class="text-2xl">Welcome!</h1>
        <p class="text-sm">Make yourself at home</p>
      </div>
    </div>
  </div>
</Layout>

<style></style>
