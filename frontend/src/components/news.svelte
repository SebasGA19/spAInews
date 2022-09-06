<script lang="ts">
	import type { Article } from '../common/types';
	import NewsEntry from './news-entry.svelte';

	export let articles: Article[] = [];

	let chunks: Article[][] = [];

	$: {
		articles;
		chunks = [];
		for (let index = 0; index < Math.floor(articles.length / 9); index++) {
			chunks.push(articles.slice(9 * index, 9 * index + 9));
		}
	}
</script>

<div class="container">
<div class="row justify-content-center mt-5 mb-5">
	{#each articles as article}
		<div class="col-md-5 col-sm-10">
			<NewsEntry {article} />
		</div>
	{/each}
</div>
</div>
