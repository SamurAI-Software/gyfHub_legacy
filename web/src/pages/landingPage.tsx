import React from "react"

// components
import LandingTop from "../components/landing/LandingTop"
import AboutSection from "../components/landing/AboutSection"
import VideoSection from "../components/landing/VideoSection"

function LandingPage() {
	return (
		<React.Fragment>
			{/* Landing */}
			<LandingTop />
			{/* About */}
			<AboutSection />
			{/* Video */}
			<VideoSection />
			{/* Footer */}
		</React.Fragment>
	)
}

export default LandingPage
