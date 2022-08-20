<script>
	import Page from '../components/page.svelte';
	import NewsSection from "../components/newsSection.svelte";
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

	loadContent();
</script>

<Page>
	<!-- By Juan Paez -->
	<header class="header">
		<h3 id="titulo">Welcome to the Jungle</h3>
	</header>
	<div>
		<!-- By Juan Paez -->
		<NewsSection bind:articles={articles}/>
		<button type="button" class="btn btn-outline-primary" on:click="{loadContent}">Ver más</button>
	</div>
</Page>

<style>
	.header {
		color: white;
		padding: 40px 0 20px;
		text-align: center;
	}

	.header h1 {
		font-size: 40px;
		font-weight: bold;
	}

	.header h2 a {
		border-bottom: 1px solid rgba(255, 255, 255, 0.5);
		color: white;
		font-size: 20px;
		opacity: 0.5;
	}

	.header h2 a:hover {
		border-bottom-color: white;
		opacity: 1;
	}
</style>