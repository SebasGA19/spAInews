<script lang="ts">
	import type { Article } from 'src/common/types';
	export let article: Article;
	export let articleModalId: string;

	function randomIntFromInterval(min: number, max: number): number {
		// min and max included
		return Math.floor(Math.random() * (max - min + 1) + min);
	}

	const articleImage: string = `/images/${article.category}-${randomIntFromInterval(1, 4)}.jpg`;
</script>

<div class="card card-blog">
	<div class="card-image">
		<img class="img" alt="..." src={articleImage} />
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
	<div class="modal-dialog modal-lg">
		<div class="modal-content">
			<div class="modal-header">
				<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close" />
			</div>
			<div class="modal-body">
				<div class="container-sm">
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
						<p class="news">
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
	/*---------------------------------------------------------------------- /
CARDS
----------------------------------------------------------------------- */
	@media screen and (width: 768px) {
		.card {
			display: inline-block;
			box-sizing: border-box;
			width: 28.5vh;
		}
	}
	@media screen and (max-width: 360px) {
		.card-description{
			font-size: 10px;
		}
		.card-caption{
			font-size: 15px;
		}
	}
	@media screen and (width:414px){
		.card-description{
			font-size: 12px;
			overflow: hidden;
			text-overflow: ellipsis;
		}
	}
	@media screen and (width: 820px) {
		.card {
			display: block;
			box-sizing: border-box;
			width: 25vh;
		}
	}
	.card {
		height: 60vh;
		display: inline-block;
		position: relative;
		margin-bottom: 40px;
		border-radius: 6px;
		box-shadow: 0 2px 2px 0 rgba(0, 0, 0, 0.14), 0 3px 1px -2px rgba(0, 0, 0, 0.2),
			0 1px 5px 0 rgba(0, 0, 0, 0.12);
	}

	.card .card-image {
		height: 50%;
		position: relative;
		overflow: hidden;
		margin-left: 2%;
		margin-right: 2%;
		margin-top: -10%;
		border-radius: 6px;
		box-shadow: 0 16px 38px -12px rgba(0, 0, 0, 0.56), 0 4px 25px 0px rgba(0, 0, 0, 0.12),
			0 8px 10px -5px rgba(0, 0, 0, 0.2);
	}

	.card .card-image img {
		border-radius: 6px;
		pointer-events: none;
		width: 100%;
		height: 100%;
		object-fit: cover;
	}
	.card .card-image .card-caption {
		position: absolute;
		bottom: 15px;
		left: 15px;
		font-size: 1.3em;
		text-shadow: 0 2px 5px rgba(33, 33, 33, 0.5);
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
	.card .ftr .author {
		bottom: 80px;
		position: absolute;
		font-size: 14px;
	}

	.card .ftr .stats {
		right: 22px;
		bottom: 70px;
		position: absolute;
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
		overflow: hidden;
		word-wrap: break-word;
		display: -webkit-box;
		-webkit-line-clamp: 3;
		-webkit-box-orient: vertical;
	}
	.card-caption {
		font-weight: 500;
		text-align: justify;
		overflow: hidden;
		display: -webkit-box;
		-webkit-line-clamp: 2;
		-webkit-box-orient: vertical;
	}
	.btn {
		bottom: 20px;
		position: absolute;
	}

	.container-sm {
		max-width: 600px;
	}

	.news {
		text-align: justify;
	}

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
		text-align: left;
		font-family: 'Times New Roman', Times, serif;
		margin-top: 10px;
		color: dodgerblue;
	}
	ul li a {
		font-size: 25px;
	}
	ul li p {
		font-size: 25px;
	}
</style>
