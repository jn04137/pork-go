<script>
  let instance = new Date()
</script>

<div class="bg-white rounded px-4 py-2 shadow">
  <div class='text-[24px] font-semibold'>
    <slot name="title">Title</slot>
  </div>
  <div class='pt-1 pb-2 text-sm h-fit'>
    <slot name="body">This is the body of the post.</slot>
  </div>
  <div class='flex flex-col text-xs'>
    <div>
      <slot name="author">Created by John Doe</slot>
    </div>
    <div class='italic'>
      <slot name="date">{instance.toLocaleString()}</slot>
    </div>
  </div>
</div>

<style>

</style>
