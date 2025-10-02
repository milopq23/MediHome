import tailwindcss from '@tailwindcss/vite';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import dotenv from 'dotenv'


dotenv.config();

const port = process.env.VITE_PORT || 3030
export default defineConfig({
	plugins: [tailwindcss(), sveltekit()],
	server:{
		port: parseInt(port)
	}
});
