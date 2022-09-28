<script lang="ts">
	import type { Article } from 'src/common/types';
	export let article: Article;
	export let articleModalId: string;
</script>

<div class="card article-card">
	<button
		style="border: none; padding: 0;"
		data-bs-toggle="modal"
		data-bs-target="#{articleModalId}"
	>
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
	</button>
</div>

<div
	class="modal"
	id={articleModalId}
	tabindex="-1"
	aria-labelledby="{articleModalId}Label"
	aria-hidden="true"
>
	<div class="modal-dialog modal-fullscreen">
		<div class="modal-content">
			<div class="modal-header">
				<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close" />
			</div>
			<div class="modal-body">
				<div class="container" style="max-width: 600px;">
					<h2 style="text-align:justify;"><a href="{article.url}" target="_blank">{article.title}</a></h2>
					<div class="row">
						<div class="col">
							<strong>Category:</strong>
							{article.category}
						</div>
						<div class="col text-end">
							<strong>News Source:</strong>
							<a href="http://{article.source_domain}">{article.source_domain}</a>
						</div>
					</div>
					<div class="row">
						<div class="col">
							<strong>Author:</strong>
							{#if article.authors === undefined || article.authors.length === 0}
								{article.source_domain}
							{:else}
								{article.authors[0]}
								{#if article.authors.length > 1}
									{#each article.authors.slice(1, article.authors.length) as author}
										{author}
									{/each}
								{/if}
							{/if}
						</div>
						<div class="col text-end">
							<strong>Date Publish:</strong>
							{new Date(article.date_publish).toLocaleDateString()}
						</div>
					</div>
					<img src="/images/{article.category}-1.jpg" class="card-img-top" alt="..." />
					<p style="text-align: justify;" class="mt-5">
						{#if article.maintext === undefined}
							{article.description}
						{:else}
							{article.maintext}
						{/if}
					</p>
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
