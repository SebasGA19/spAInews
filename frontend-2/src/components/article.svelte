<script>
	export let image_url;
	export let title;
	export let description;
	export let authors;
	export let publish_date;
</script>

<a class="article first-article" href="#">
	<figure class="article-image is-4by3">
		<img src={image_url} alt="" />
	</figure>
	<div class="article-body">
		<h2 class="article-title">
			{title}
		</h2>
		<p class="article-content">
			{description}
			<a class="article-news" href="#">Continuar leyendo</a>
		</p>
		<footer class="article-info">
			{#each authors as author}
				<span>{author}</span>
			{/each}
			<span class="date-published">{publish_date}</span>
		</footer>
	</div>
</a>

<style>
	
</style>
