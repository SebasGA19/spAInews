<script>
	import Page from '../components/page.svelte';
	import NewsSection from "../components/newsSection.svelte";
	import { onMount } from 'svelte';
	import { latestnews } from '../api/latestnews';

	let currentPage = 1;
	let articles = [];

	function loadContent() {
		latestnews(currentPage).then(
			(results) => {
				currentPage += 1;
				articles = articles.concat(results.news);
			}
		);
	}

	onMount(
		async () => {
			loadContent();
		}
	);
</script>

<Page>
	<!-- By Juan Paez -->
	<header class="header">
		<h3 id="titulo">Welcome to the Jungle</h3>
	</header>
	<div>
		<!-- By Juan Paez -->
		<NewsSection bind:articles={articles}/>
		<button type="button" class="btn btn-outline-primary">Ver más</button>
	</div>
</Page>
