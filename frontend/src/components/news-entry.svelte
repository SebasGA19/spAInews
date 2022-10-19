<script lang="ts">
	import type { Article } from 'src/common/types';
	export let article: Article;
	export let articleModalId: string;

	function randomIntFromInterval(min: number, max: number): number {
		// min and max included
		return Math.floor(Math.random() * (max - min + 1) + min);
	}

	const articleImage: string = `/images/${article.category}/${randomIntFromInterval(1, 4)}.jpg`;
</script>

<div class="card w-100 article-card">
	<div class="article-image">
		<img class="img" alt="..." src={articleImage} />
	</div>
	<div class="p-4 article-card-body">
		<h6 class="mt-0 text-info mt-0 mb-0">{article.category}</h6>
		<div class="article-card-details-block">
			<h4 class="card-title">
				{article.title}
			</h4>
			<p class="card-description">
				{article.description}
			</p>
		</div>
		<div class="row w-100 mt-0">
			<div class="col article-card-author">
				<strong>
					{#if article.authors === undefined || article.authors.length === 0}
						{article.source_domain}
					{:else}
						{article.authors[0]}
						{#if article.authors.length > 1}
							{#each article.authors.slice(1, article.authors.length) as author}
								, {author}
							{/each}
						{/if}
					{/if}
				</strong>
			</div>
			<div class="col w-100 d-flex justify-content-end">
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
	<div class="modal-dialog modal-lg">
		<div class="modal-content">
			<div class="modal-header">
				<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close" />
			</div>
			<div class="modal-body">
				<div class="container" style="max-width: 600px;">
					<div class="heading">
						<h1>
							<a style="text-decoration: none;" href={article.url} target="_blank"
								>{article.title}</a
							>
						</h1>
						<img style="width:100%; border-radius: 5px;" src={articleImage} alt="fotito" />
						<br />
						<p style="font-size:30px;" class="modal-category">
							{article.category}
						</p>
						<ul>
							{#if article.authors !== undefined && article.authors.length >= 1}
								<li>
									<p>
										{article.authors[0]}
										{#if article.authors.length > 1}
											{#each article.authors.slice(1, article.authors.length) as author}
												, {author}
											{/each}
										{/if}
									</p>
								</li>
								<li><p>|</p></li>
							{/if}
							<li><p>{new Date(article.date_publish).toLocaleDateString()}</p></li>
							<li><p>|</p></li>
							<li>
								<a style="text-decoration:none;" href="http://{article.source_domain}"
									>{article.source_domain}</a
								>
							</li>
						</ul>
					</div>

					<div>
						<hr />
						<p class="text-justify">
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
</div>

<style>
	/*
	Card
	*/
	.article-card {
		border-radius: 6px;
		box-shadow: 0 2px 2px 0 rgba(0, 0, 0, 0.14), 0 3px 1px -2px rgba(0, 0, 0, 0.2),
			0 1px 5px 0 rgba(0, 0, 0, 0.12);
	}

	.article-image {
		transition: 1s;
		height: 350px;
		margin-left: 15px;
		margin-right: 15px;
		margin-top: -30px;
		border-radius: 6px;
		box-shadow: 0 16px 10px -12px rgba(0, 0, 0, 0.56), 0 4px 25px 0px rgba(0, 0, 0, 0.12),
			0 8px 10px -5px rgba(0, 0, 0, 0.2);
	}

	.article-image:hover {
		transform: scale(1.02);
	}

	.article-image img {
		object-fit: cover;
		width: 100%;
		height: 100%;
		border-radius: 6px;
	}

	.article-card-details-block {
		height: 210px;
		width: 100%;
	}
	.card-title {
		max-height: 105px;
		text-align: justify;
		overflow: hidden;
		line-height: normal;
		display: -webkit-box;
		-webkit-box-orient: vertical;
		-webkit-line-clamp: 3;
		text-overflow: ellipsis;
	}

	.card-description {
		max-height: 105px;
		text-align: justify;
		overflow: hidden;
		line-height: normal;
		display: -webkit-box;
		-webkit-box-orient: vertical;
		-webkit-line-clamp: 3;
		text-overflow: ellipsis;
	}

	.article-card-author {
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis;
	}

	ul li a,
	p {
		font-size: large;
	}

	/*
	Modal
	*/

	.heading ul {
		list-style: none;
		margin: 0;
		display: inline-block;
		padding: 0;
	}
	.heading li {
		padding: 5px;
		display: block;
		float: left;
		font-size: 12px;
	}

	.modal-category {
		margin-top: 10px;
		font-size: 18px;
		color: dodgerblue;
		text-align: left;
		font-family: 'Times New Roman', Times, serif;
	}
</style>
