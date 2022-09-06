<script lang="ts">
	import { updatePassword } from '../../api/change-password';
	import { updateEmail } from '../../api/change-email';
	import { session } from '../../stores/session';
	import { accountDetails } from '../../api/account';
	import AuthRequired from '../../components/auth-required.svelte';
	import { onMount } from 'svelte';
	import type { Account, Words } from '../../common/types';
	import { getWords, updateWords } from '../../api/words';

	let userAccount: Account = { username: '', email: '', 'creation-date': '' };
	let userWords: Words = { automatic: false, words: [] };

	let emailFormPassword: string;
	let emailFormNewEmail: string;
	let passwordFormOldPassword: string;
	let passwordFormNewPassword: string;
	let emailFormErrorMessage: string | undefined;
	let emailFormSuccessMessage: string | undefined;
	let passwordFormErrorMessage: string | undefined;
	let passwordFormSuccessMessage: string | undefined;
	let wordsFormNewWord: string;
	let wordsFormErrorMessage: string | undefined;
	let wordsFormSuccessMessage: string | undefined;

	onMount(async () => {
		accountDetails($session)
			.then((account) => {
				userAccount = account;
			})
			.catch((error) => {
				alert(error);
			});
		getWords($session)
			.then((words) => {
				userWords = words;
			})
			.catch((error) => {
				alert(error);
			});
	});

	function handleChangePassword() {
		updatePassword(passwordFormOldPassword, passwordFormNewPassword, $session)
			.then(() => {
				passwordFormErrorMessage = undefined;
				passwordFormSuccessMessage = 'Password updated';
			})
			.catch((error) => {
				passwordFormSuccessMessage = undefined;
				passwordFormErrorMessage = error;
			});
	}

	function handleChangeEmail() {
		updateEmail(emailFormPassword, emailFormNewEmail, $session)
			.then(() => {
				emailFormErrorMessage = undefined;
				emailFormSuccessMessage = 'Check your email for authentication';
			})
			.catch((error) => {
				emailFormSuccessMessage = undefined;
				emailFormErrorMessage = error;
			});
	}

	function handleUpdateWords() {
		if (userWords.automatic) {
			userWords.words = [];
		}
		updateWords(userWords, $session)
			.then(() => {
				wordsFormErrorMessage = undefined;
				wordsFormSuccessMessage = 'Words updated';
			})
			.catch((error) => {
				wordsFormSuccessMessage = undefined;
				wordsFormErrorMessage = error;
			});
	}

	function addWord() {
		userWords.words.push(wordsFormNewWord);
		userWords.words = (Array<string>).from(new Set<string>(userWords.words));
		handleUpdateWords();
	}

	function deleteFromWords(index: number): () => void {
		return () => {
			userWords.words.splice(index, 1);
			userWords.words = userWords.words;
			handleUpdateWords();
		};
	}
</script>

<AuthRequired />
<div class="container">
	<h3>Welcome <strong>{userAccount.username}</strong>!</h3>
	<div class="accordion accordion-flush" id="settingsAccordionFlush">
		<div class="accordion-item">
			<h2 class="accordion-header" id="flush-email">
				<button
					class="accordion-button collapsed"
					type="button"
					data-bs-toggle="collapse"
					data-bs-target="#flush-collapseEmail"
					aria-expanded="false"
					aria-controls="flush-collapseEmail"
				>
					Change Email
				</button>
			</h2>
			<div
				id="flush-collapseEmail"
				class="accordion-collapse collapse"
				aria-labelledby="flush-email"
				data-bs-parent="#settingsAccordionFlush"
			>
				<div class="accordion-body">
					<div class="container" style="max-width: 500px;">
						{#if emailFormErrorMessage}
							<div class="alert alert-danger" role="alert">
								{emailFormErrorMessage}
							</div>
						{/if}
						{#if emailFormSuccessMessage}
							<div class="alert alert-success" role="alert">
								{emailFormSuccessMessage}
							</div>
						{/if}
						<form on:submit|preventDefault={handleChangeEmail}>
							<div class="mb-3">
								<label for="password" class="form-label">Password</label>
								<input
									required
									type="password"
									placeholder="password"
									bind:value={emailFormPassword}
									class="form-control"
								/>
								<label for="change-email" class="form-label">New email</label>
								<input
									required
									type="email"
									placeholder={userAccount.email}
									bind:value={emailFormNewEmail}
									class="form-control"
								/>
							</div>
							<button type="submit" class="btn btn-primary">Submit</button>
						</form>
					</div>
				</div>
			</div>
		</div>
		<div class="accordion-item">
			<h2 class="accordion-header" id="flush-headingTwo">
				<button
					class="accordion-button collapsed"
					type="button"
					data-bs-toggle="collapse"
					data-bs-target="#flush-collapsePassword"
					aria-expanded="false"
					aria-controls="flush-collapsePassword"
				>
					Change Password
				</button>
			</h2>
			<div
				id="flush-collapsePassword"
				class="accordion-collapse collapse"
				aria-labelledby="flush-headingTwo"
				data-bs-parent="#settingsAccordionFlush"
			>
				<div class="accordion-body">
					<div class="container" style="max-width: 500px;">
						{#if passwordFormErrorMessage}
							<div class="alert alert-danger" role="alert">
								{passwordFormErrorMessage}
							</div>
						{/if}
						{#if passwordFormSuccessMessage}
							<div class="alert alert-success" role="alert">
								{passwordFormSuccessMessage}
							</div>
						{/if}
						<form on:submit|preventDefault={handleChangePassword}>
							<div class="mb-3">
								<label for="change-password" class="form-label">Old password</label>
								<input
									required
									type="password"
									placeholder="old-password"
									bind:value={passwordFormOldPassword}
									class="form-control"
								/>
								<label for="change-password" class="form-label">New password</label>
								<input
									required
									type="password"
									placeholder="new-password"
									bind:value={passwordFormNewPassword}
									class="form-control"
								/>
							</div>
							<button type="submit" class="btn btn-primary">Submit</button>
						</form>
					</div>
				</div>
			</div>
		</div>
		<div class="accordion-item">
			<h2 class="accordion-header" id="flush-headingTwo">
				<button
					class="accordion-button collapsed"
					type="button"
					data-bs-toggle="collapse"
					data-bs-target="#flush-collapseWords"
					aria-expanded="false"
					aria-controls="flush-collapseWords"
				>
					Words
				</button>
			</h2>
			<div
				id="flush-collapseWords"
				class="accordion-collapse collapse"
				aria-labelledby="flush-headingTwo"
				data-bs-parent="#settingsAccordionFlush"
			>
				<div class="accordion-body">
					<div class="container" style="max-width: 500px;">
						{#if wordsFormErrorMessage}
							<div class="alert alert-danger" role="alert">
								{wordsFormErrorMessage}
							</div>
						{/if}
						{#if wordsFormSuccessMessage}
							<div class="alert alert-success" role="alert">
								{wordsFormSuccessMessage}
							</div>
						{/if}
						<div class="form-check form-switch mb-3">
							<input
								class="form-check-input"
								type="checkbox"
								role="switch"
								id="wordsAutomatic"
								bind:checked={userWords.automatic}
								on:change={handleUpdateWords}
								value={userWords.automatic}
							/>
							<label class="form-check-label" for="wordsAutomatic">Pick words automatic</label>
						</div>
						{#if !userWords.automatic}
							<div class="input-group mb-3">
								<input
									type="text"
									class="form-control"
									placeholder="finance"
									aria-label="finance"
									aria-describedby="button-add-word"
									bind:value={wordsFormNewWord}
								/>
								<button
									class="btn btn-primary"
									type="button"
									id="button-add-word"
									on:click={addWord}>Add</button
								>
							</div>
							<div class="justify-content-center text-center">
								{#each userWords.words as word, index}
									<div class="row mb-3">
										<div class="col-7">
											<h6 style="overflow-x: hidden; overflow-y: hidden;">{word}</h6>
										</div>
										<div class="col-5">
											<button type="button" class="btn btn-danger" on:click={deleteFromWords(index)}
												>Delete</button
											>
										</div>
									</div>
								{/each}
							</div>
						{/if}
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
