<template>
    <header class="bg-dark p-3 text-white">
        <div class="container">
            <div class="d-flex flex-wrap align-items-center justify-content-center justify-content-lg-start">
                <ul class="nav col-12 col-lg-auto me-lg-auto mb-2 justify-content-center mb-md-0">
                    <li>
                        <router-link class="nav-link px-3 text-secondary" to="/">Home</router-link>
                    </li>
                    <li>
                        <router-link class="nav-link px-3 text-white" to="/projects">Projects</router-link>
                    </li>
                    <li>
                        <router-link class="nav-link px-3 text-secondary" to="/cv">CV</router-link>
                    </li>
                </ul>
            </div>
        </div>
    </header>
</template>