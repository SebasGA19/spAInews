<script type="ts">
	import { onMount } from 'svelte';

	let seconds: number = 5;
	onMount(() => {
		setInterval(() => {
			if (seconds > 0) seconds--;
		}, 1000);
	});

	$ : {
		seconds;
		if (seconds <= 0) {
			window.location.href = "/";
		}
	}
</script>

<div class="text-center">
	<h2>Registration completed</h2>
	<h4>Automatically redirecting to home in {seconds} {seconds !== 1 ? "seconds" : "second"}</h4>
	<a href="/" class="btn btn-primary">Go home</a>
</div>
