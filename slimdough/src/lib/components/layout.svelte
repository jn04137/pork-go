<script>
  import Navbar from './navbar.svelte'
  import Footer from './footer.svelte'

  let width = 1000
</script>

<div class='h-screen flex flex-col justify-between'>
  <Navbar/>
  <div class='flex grow justify-center content'>
    <div class="flex min-w-[1000px] py-4">
      <slot name='content'>This is the body</slot>
    </div>
  </div>
  <div class='flex justify-center py-2'>
    <div class="min-w-[1000px]">
      <Footer/>
    </div>
  </div>
</div>

<style>
.content {
  background-color: #F2F2F2
}
</style>
