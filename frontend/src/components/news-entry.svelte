<script lang="ts">
	import type { Article } from 'src/common/types';
	export let article: Article;
	console.log(article)
</script>

<div class="card article-card">
	<img src="/images/{article.category}-1.jpg" class="card-img-top" alt="..." />
	<div class="card-body">
		<p class="card-text">
			{article.description}
		</p>
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
