<script type="ts">
	import NewsSection from '../../components/news.svelte';
	import { search } from '../../api/news';
	import type { Article, SearchFilter } from 'src/common/types';

	let currentPage: number = 0;
	let articles: Article[] = [];

    let fromDate: string;
    let toDate: string;
    let titleSearchQuery: string;
	let maintextSearchQuery: string;
	let searchFilter: SearchFilter;

	function loadMore() {
		search(currentPage, searchFilter)
			.then((results) => {
				if (results['number-of-results'] > 0) {
					currentPage += 1;
					articles = articles.concat(results.news);
				}
			})
			.catch((error) => {
				alert(error);
			});
	}

	function doSearch(): void {
        fromDate = '';
        toDate = '';
        titleSearchQuery = '';
		articles = [];
		currentPage = 1;
		searchFilter = {
			'maintext-words': maintextSearchQuery.split(' ')
		};

		loadMore();
	}
    
    function updateFilter() {
        articles = [];
		currentPage = 1;
        searchFilter = {
			'maintext-words': maintextSearchQuery.split(' ')
		};
        if (fromDate !== '') {
            searchFilter['start-date'] = new Date(fromDate).toISOString();
        }
        if (toDate !== '') {
            searchFilter['end-date'] = new Date(toDate).toISOString();
        }
        if (titleSearchQuery !== '') {
            searchFilter['title-words'] = titleSearchQuery.split(' ')
        }
        loadMore();
    }
</script>

<div class="container text-center" style="max-width: 700px;">
	<form on:submit|preventDefault={doSearch}>
		<div class="d-flex border">
			<input
				type="text"
				required
				placeholder="Search"
				class="form-control"
				style="border: none;"
				bind:value={maintextSearchQuery}
			/>
			{#if searchFilter !== undefined}
				<button
					class="btn text-reset"
					style="background-color: transparent;"
					type="button"
					data-bs-toggle="offcanvas"
					data-bs-target="#filterMenu"
					aria-controls="filterMenu"
				>
					<i class="fa-solid fa-filter" />
				</button>
			{/if}
		</div>
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

<!--Filter-->
<div
	class="offcanvas offcanvas-end"
	tabindex="-1"
	id="filterMenu"
	aria-labelledby="filterMenuLabel"
>
	<div class="offcanvas-header">
		<h5 class="offcanvas-title" id="filterMenuLabel">Advance filter</h5>
		<button type="button" class="btn-close" data-bs-dismiss="offcanvas" aria-label="Close" />
	</div>
	<div class="offcanvas-body">
        <input
				type="text"
				placeholder="Search inside title"
				class="form-control"
				bind:value={titleSearchQuery}
                on:change={updateFilter}
			/>
        <label>From date:</label>
		<input type="date" class="form-control" bind:value={fromDate} on:change={updateFilter}>
        <label>To date:</label>
		<input type="date" class="form-control" bind:value={toDate} on:change={updateFilter}>
	</div>
</div>
