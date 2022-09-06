<script lang="ts">
	import type { Article } from "src/common/types";
	export let article: Article;
</script>

<div class="card">
	{JSON.stringify(article)}
</div>

