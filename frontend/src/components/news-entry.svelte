<script lang="ts">
	import type { Article } from 'src/common/types';
	export let article: Article;
	export let articleModalId: string;
</script>
<div class="card card-blog">
	<div class="card-image">
		<img class="img" alt="..." src="/images/{article.category}-1.jpg" />
		<div class="ripple-cont" />
	</div>
	<div class="table">
		<h6 class="category text-info">{article.category}</h6>
		<h3 class="card-caption">
			{article.title}
		</h3>
		<p class="card-description">
			{article.description}
		</p>
		<div class="ftr">
			<div class="author">
				<strong>
					{#if article.authors === undefined || article.authors.length === 0}
						{article.source_domain}
					{:else}
						{article.authors[0]}
						{#if article.authors.length > 1}
							...
						{/if}
					{/if}
				</strong>
			</div>
			<div class="stats">
				<strong> {new Date(article.date_publish).toLocaleDateString()}</strong>
			</div>
		</div>
		<br />
		<button class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#{articleModalId}"
			>Read more</button
		>
	</div>
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
					<h2 style="text-align:justify;">
						<a href={article.url} target="_blank">{article.title}</a>
					</h2>
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
/*---------------------------------------------------------------------- /
CARDS
----------------------------------------------------------------------- */
	.card {
		display: inline-block;
		position: relative;
		width: 100%;
		margin-bottom: 40px;
		border-radius: 6px;
		box-shadow: 0 2px 2px 0 rgba(0, 0, 0, 0.14), 0 3px 1px -2px rgba(0, 0, 0, 0.2),
			0 1px 5px 0 rgba(0, 0, 0, 0.12);
	}

	.card .card-image {
		height: 60%;
		position: relative;
		overflow: hidden;
		margin-left: 15px;
		margin-right: 15px;
		margin-top: -30px;
		border-radius: 6px;
		box-shadow: 0 16px 38px -12px rgba(0, 0, 0, 0.56), 0 4px 25px 0px rgba(0, 0, 0, 0.12),
			0 8px 10px -5px rgba(0, 0, 0, 0.2);
	}

	.card .card-image img {
		width: 100%;
		height: 100%;
		border-radius: 6px;
		pointer-events: none;
	}

	.card .card-image .card-caption {
		position: absolute;
		bottom: 15px;
		left: 15px;
		font-size: 1.3em;
		text-shadow: 0 2px 5px rgba(33, 33, 33, 0.5);
	}

	.card img {
		width: 100%;
		height: auto;
	}

	.card .ftr {
		margin-top: 15px;
	}

	.card .ftr div {
		display: inline-block;
	}

	.card .ftr .stats {
		float: right;
		line-height: 30px;
	}

	.card .ftr .stats {
		position: relative;
		top: 1px;
		font-size: 14px;
	}

	/* ============ Card Table ============ */

	.table {
		margin-bottom: 0px;
	}

	.card .table {
		padding: 15px 30px;
	}

	/* ============ Card Blog ============ */

	.card-blog {
		margin-top: 30px;
	}

	/* ============ Card Category ============ */

	.category {
		position: relative;
		line-height: 0;
		margin: 15px 0;
	}

	/* ============ Card Author ============ */

	.card .author a .ripple-cont {
		display: none;
	}

	/* ============ Card Product ============ */

	.card-product .ftr .stats {
		margin-top: 4px;
		top: 0;
	}

	.card-description {
		text-align: justify;
	}

	.card-caption {
		font-weight: 700;
		text-align: justify;
	}
</style>
