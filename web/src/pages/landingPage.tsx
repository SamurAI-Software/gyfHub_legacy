import React from "react"

// components
import LandingTop from "../components/landing/LandingTop"
import AboutSection from "../components/landing/AboutSection"
import VideoSection from "../components/landing/VideoSection"
import GifGrid from "../components/landing/GifGrid"
import Footer from "../components/landing/Footer"

function LandingPage() {
	return (
		<React.Fragment>
			{/* Landing */}
			<LandingTop />
			{/* About */}
			<AboutSection />
			{/* Gif Grid */}
			<GifGrid />
			{/* Video */}
			<VideoSection />
			{/* Footer */}
			<Footer />
		</React.Fragment>
	)
}

export default LandingPage
