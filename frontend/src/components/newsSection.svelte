<script>
	import Article from './article.svelte';

	export let articles;

	let chunks = [];

	$: {
		articles;
        chunks = [];
		for (let index = 0; index < parseInt(articles.length / 9); index++) {
			chunks.push(articles.slice(9 * index, 9 * index + 9));
		}
	}
</script>

{#key articles}
	{#each chunks as chunk}
		<main class="main columns">
			<section class="column main-column">
				<a class="article first-article" href="#">
					{#if chunk.length > 0}
						<Article
							image_url={chunk[0].image_url}
							title={chunk[0].title}
							description={chunk[0].description}
							authors={chunk[0].authors}
							publish_date={chunk[0].date_publish}
						/>
					{/if}
				</a>

				<div class="columns">
					<div class="column nested-column">
						{#if chunk.length > 1}
							<Article
								image_url={chunk[1].image_url}
								title={chunk[1].title}
								description={chunk[1].description}
								authors={chunk[1].authors}
								publish_date={chunk[1].date_publish}
							/>
						{/if}
					</div>

					<div class="column">
						{#if chunk.length > 2}
							{#each chunk.slice(2, 5) as article}
								<Article
									title={article.title}
									description={article.description}
									authors={article.authors}
									publish_date={article.date_publish}
								/>
							{/each}
						{/if}
					</div>
				</div>
			</section>

			<section class="column">
				{#if chunk.length > 5}
                    <Article
                        image_url={chunk[5].image_url}
					    title={chunk[5].title}
					    description={chunk[5].description}
					    authors={chunk[5].authors}
					    publish_date={chunk[5].date_publish}
					/>
					{#each chunk.slice(6, 10) as article}
						<Article
							title={article.title}
							description={article.description}
							authors={article.authors}
							publish_date={article.date_publish}
						/>
					{/each}
				{/if}
			</section>
		</main>
	{/each}
{/key}

<style>
	.main {
		margin: 0 auto;
		max-width: 1000px;
		padding: 10px;
	}

	.column {
		flex: 1;
		flex-direction: column;
	}

	@media screen and (min-width: 800px) {
		.columns,
		.column {
			display: flex;
		}
	}
	@media screen and (min-width: 1000px) {
		.first-article {
			flex-direction: row;
		}

		.first-article .article-body {
			flex: 1;
		}

		.first-article .article-image {
			height: 300px;
			order: 2;
			padding-top: 0;
			width: 400px;
		}

		.main-column {
			flex: 3;
		}

		.nested-column {
			flex: 2;
		}
	}
</style>
