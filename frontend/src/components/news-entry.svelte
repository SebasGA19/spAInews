<script lang="ts">
	import type { Article } from 'src/common/types';
	export let article: Article;
</script>

<div class="card article-card">
	<img src="/images/{article.category}-1.jpg" class="card-img-top" alt="..." />
	<div class="card-body">
		<h3 style="text-align:justify;">{article.title}</h3>
		<div class="container">
			<p class="card-text" style="font-size: large; text-align: justify;">
				{article.description}
			</p>
		</div>

		<div class="container">
			<div class="row">
				<div class="col">
					<strong>Author:</strong>
					{#if article.authors === undefined || article.authors.length === 0}
						{article.source_domain}
					{:else}
						{article.authors[0]}
						{#if article.authors.length > 1}
							...
						{/if}
					{/if}
				</div>

				<div class="col text-end">
					<strong>Date Publish:</strong>
					{new Date(article.date_publish).toLocaleDateString()}
				</div>
			</div>
		</div>
	</div>
</div>

<style>
	.article-card {
		transition: 0.5s;
	}
	.article-card:hover {
		scale: 1.04;
	}
</style>
