<script>
	import { alerts } from '../stores/alerts.js';
</script>

<div class="alert-container">
	{#each $alerts as alert (alert.id)}
		<div class="alert {alert.type}">
			{alert.message}
		</div>
	{/each}
</div>

<style>
	.alert-container {
		position: fixed;
		top: 1rem;
		right: 1rem;
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
		z-index: 999;
	}
	.alert {
		padding: 0.5rem 1rem;
		border-radius: 4px;
		color: white;
		box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
	}
	.alert.success {
		background-color: green;
	}
	.alert.error {
		background-color: red;
	}
</style>
