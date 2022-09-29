<script type="ts">
	import NewsSection from '../../components/news.svelte';
	import { search } from '../../api/news';
	import type { Article, SearchFilter } from 'src/common/types';
	import { onMount } from 'svelte';

	let currentPage: number = 0;
	let articles: Article[] = [];

    let searchQuery: string;
    let searchFilter: SearchFilter = {};

	function doSearch(): void {
        articles = [];
        currentPage = 1;
        searchFilter = {
            "maintext-words": searchQuery.split(" ")
        };
		search(currentPage, searchFilter)
			.then((results) => {
				if (results['number-of-results'] > 0) {
					currentPage += 1;
					articles = articles.concat(results.news);
				}
			})
			.catch((error) => {alert(error)});
	}
    function loadMore() {
        search(currentPage, searchFilter)
			.then((results) => {
				if (results['number-of-results'] > 0) {
					currentPage += 1;
					articles = articles.concat(results.news);
				}
			})
			.catch((error) => {alert(error)});
    }
</script>

<div class="container text-center" style="max-width: 700px;">
	<form on:submit|preventDefault={doSearch}>
        <input type="text" required placeholder="Search" class="form-control" bind:value={searchQuery}>
    </form>
</div>
<div>
	<NewsSection bind:articles />
    {#if currentPage > 0}
	<div class="text-center">
		<button type="button" class="btn btn-primary" on:click={loadMore}>More</button>
	</div>
    {/if}
</div>
