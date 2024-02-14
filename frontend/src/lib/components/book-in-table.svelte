<script>
	export let book = {};
	export let admin = false;
</script>

<details>
	<summary>{book.name}</summary>
	<ul>
		<li><span class="fw-bold">Gênero: </span>{book.gender}</li>
		<li><span class="fw-bold">Exemplares disponíveis: </span>{book.copies}</li>
		<li>
			<span class="fw-bold">Descrição: </span>
			<p>{book.description}</p>
		</li>
	</ul>
	<div class="d-flex justify-content-end">
		<slot />
	</div>
</details>
