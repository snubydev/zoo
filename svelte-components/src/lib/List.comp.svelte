<svelte:options customElement={{
  tag: 'zoo-list',
  shadow: 'none',
  props: {
    custom_items: { reflect: true, type: 'Array' },
    title: { reflect: true, type: 'String' },
  },
}} />

<script>
  export let custom_items = []
  export let title = 'Items list'
</script>

{#if custom_items && custom_items.length > 0}
  <div class="my-4 text-2xl">
    {title}
  </div>
{/if}

{#each custom_items || [] as p}
  <div class="mb-2 text-lg">{p}</div>
{/each}
