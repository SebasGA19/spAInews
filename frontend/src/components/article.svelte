<script>
	export let image_url = undefined;
	export let title;
	export let description;
	export let authors;
	export let publish_date;
</script>

<a class="article first-article" href="#">
    {#if image_url !== undefined}
	<figure class="article-image is-4by3">
		<img src={image_url} alt="" />
	</figure>
    {/if}
	<div class="article-body">
		<h2 class="article-title">
			{title}
		</h2>
		<p class="article-content">
			{description}
			<a class="article-news" href="#">Continuar leyendo</a>
		</p>
		<footer class="article-info">
			{#if authors.length > 1}
				<span>{authors[0]}</span>
			{/if}
			<span class="date-published">{publish_date}</span>
		</footer>
	</div>
</a>

<style>
	.article {
		background: white;
		box-shadow: 0 1px 3px rgba(0, 0, 0, 0.12), 0 1px 2px rgba(0, 0, 0, 0.24);
		color: #666;
		display: flex;
		flex: 1;
		flex-direction: column;
		flex-basis: auto;
		margin: 10px;
		transition: 0.5s;
	}

	.article:hover,
	.article:focus {
		transform: scale(1.02);
	}

	.article-image {
		background: #eee;
		display: block;
		padding-top: 75%;
		position: relative;
		width: 100%;
		height: 100%;
	}

	.article-image img {
		display: block;
		height: 100%;
		left: 0;
		position: absolute;
		top: 0;
		width: 100%;
	}

	.article-image.is-3by2 {
		padding-top: 66.6666%;
	}

	.article-image.is-16by9 {
		padding-top: 56.25%;
	}

	.article-body {
		display: flex;
		flex: 1;
		flex-direction: column;
		padding: 20px;
	}

	.article-title {
		flex-shrink: 0;
		font-size: 1.4em;
		font-weight: 700;
		line-height: 1.2;
	}

	.article-content {
		flex: 1;
		margin-top: 5px;
	}

	.article-info {
		display: flex;
		font-size: 0.85em;
		justify-content: space-between;
		margin-top: 10px;
	}
	.article-news {
		background: linear-gradient(black, white);
		-webkit-background-clip: text;
		-webkit-text-fill-color: transparent;
		font-size: 18px;
	}
	.article-news :hover {
		border-bottom: 1px solid;
		background: gray;
	}
</style>
