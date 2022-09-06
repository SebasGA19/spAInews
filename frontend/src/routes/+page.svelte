<script lang="ts">
	import NewsSection from '../components/news.svelte';
	import { news } from '../api/news';
	import type { Article } from 'src/common/types';
	import { onMount } from 'svelte';

	let currentPage: number = 1;
	let articles: Article[] = [];

	function loadContent() {
		news(currentPage)
			.then((results) => {
				currentPage += 1;
				articles = articles.concat(results.news);
			})
			.catch((error) => {});
	}

	onMount(loadContent);
</script>

<div class="text-center">
	<h3>Latest news</h3>
</div>
<div>
	<NewsSection bind:articles />
	<div class="text-center">
		<button type="button" class="btn btn-primary" on:click={loadContent}>More</button>
	</div>
</div>
