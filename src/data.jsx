import React from "react";
import {
  BiServer,
  BiNews,
  BiDetail,
  BiRocket,
  BiPaperPlane,
  BiHelpCircle,
  BiLogIn,
  BiUser,
  BiError,
} from "react-icons/bi";
import { HiOutlineAcademicCap } from "react-icons/hi";
import { BsHeadset } from "react-icons/bs";
import { FaLaptopCode, FaRegClock, FaRegTrashAlt } from "react-icons/fa";

const navServiceLinks = [
  {
    title: "Data Analytics",
    href: "/services/data-analytics",
    icon: <i className="fas fa-chart-line"></i>,
  },
  {
    title: "Cyber Security",
    href: "/services/cyber-security",
    icon: <i className="fas fa-shield-alt"></i>,
  },
  {
    title: "Cloud Services",
    href: "/services/cloud-services",
    icon: <i className="fas fa-cloud"></i>,
  },
  {
    title: "Applied AI",
    href: "/services/applied-ai",
    icon: <i className="fas fa-brain"></i>,
  },
  {
    title: "Security testing",
    href: "/services/security-testing",
    icon: <i className="fas fa-bug"></i>,
  },
  {
    title: "Managed Services",
    href: "/services/managed-services",
    icon: <i className="fas fa-shield-alt"></i>,
  },
];

const navInsightsLinks = [
  {
    title: "Blog",
    href: "/BlogPage",
    icon: <i className="fas fa-blog"></i>,
  },
  {
    title: "Newsletter",
    href: "/insights/newsletter",
    icon: <i className="fas fa-envelope"></i>,
  },
  {
    title: "Care Study",
    href: "/insights/case-study",
    icon: <i className="fas fa-book"></i>,
  },
];

const navProductsLinks = [
  {
    title: "PramitiHR",
    href: "#pramitihr",
    icon: <i className="fas fa-code"></i>,
  },
  {
    title: "Lanquill",
    href: "#lanquill",
    icon: <i className="fas fa-code"></i>,
  },
  {
    title: "Gen AI",
    href: "#genAI",
    icon: <i className="fas fa-robot"></i>,
  },
];

const navCompanyLinks = [
  {
    title: "Contact Us",
    icon: <BiLogIn />,
    href: "/contact-us",
  },
  {
    title: "Service Single",
    icon: <BiServer />,
    href: "/single-service",
  },
  {
    title: "Our Latest News",
    icon: <BiNews />,
    href: "/blogs",
  },
  {
    title: "News Details",
    icon: <BiDetail />,
    href: "/blog-single",
  },
  {
    title: "Career",
    icon: <HiOutlineAcademicCap />,
    href: "/career",
  },
  {
    title: "Career Single",
    icon: <HiOutlineAcademicCap />,
    href: "/career-single",
  },
  {
    title: "Integrations",
    icon: <BiRocket />,
    href: "/integrations",
  },
  {
    title: "Integrations Single",
    icon: <BiPaperPlane />,
    href: "/integration-single",
  },
];

const navCompanyPage = [
  {
    title: "Help Center",
    icon: <BiHelpCircle />,
    href: "/help-center",
  },
  {
    title: "Help Details",
    icon: <BsHeadset />,
    href: "/help-center-single",
  },
  {
    title: "Request for Demo",
    icon: <FaLaptopCode />,
    href: "/request-demo",
  },
  {
    title: "User Login",
    icon: <BiLogIn />,
    href: "/login",
  },
  {
    title: "User SignUp",
    icon: <BiUser />,
    href: "/register",
  },
  {
    title: "Recovery Account",
    icon: <FaRegTrashAlt />,
    href: "/password-reset",
  },
  {
    title: "404 Page",
    icon: <BiError />,
    href: "/404",
  },
  {
    title: "Coming Soon",
    icon: <FaRegClock />,
    href: "/coming-soon",
  },
];

const offcanvasMenuData = [
  {
    title: "Sass Company 1",
    info: "It's for SaaS Software Company",
    href: "/",
  },
  {
    title: "Sass Company 2",
    info: " Modern Sass agency",
    href: "/home-sass-two",
  },
  {
    title: "Desktop App",
    info: "Web Software Company",
    href: "/home-desktop-app",
  },

  {
    title: "App Landing",
    info: " App and Software Landing",
    href: "/home-app-landing",
  },
  {
    title: "Software Application",
    info: "IT solutions and SaaS Application",
    href: "/home-soft-app",
  },
  {
    title: "Startup Agency",
    info: "Different type of Agency",
    href: "/home-startup",
  },
  {
    title: "Data Analysis",
    info: " Software & Data Analysis",
    href: "/home-data-analysis",
  },
  {
    title: "App Landing Two",
    info: "Software & Data Analysis",
    href: "/app-landing-two",
  },
  {
    title: "IT Solution",
    info: "IT solution and sass application",
    href: "/home-nine",
  },
  {
    title: "Cyber Security",
    info: "Cyber security landing page",
    href: "/home-ten",
  },
  {
    title: "Crypto Currency",
    info: "Crypto Currency landing page",
    href: "/home-eleven",
  },
  {
    title: "Game Solutions",
    info: "Game Server landing page",
    href: "/home-twelve",
  },
  {
    title: "Payment Gateway",
    info: "Payment Gateway landing page",
    href: "/home-thirteen",
  },
  {
    title: "Digital Marketing",
    info: "Digital Marketing landing page",
    href: "/home-fourteen",
  },
  {
    title: "Conference & Event",
    info: "Conference & Event landing",
    href: "/home-fifteen",
  },

  {
    title: "Quiety Insurance",
    info: "Quiety Landing Page",
    href: "/quiety-insurance",
  },

  {
    title: "Sass Marketing",
    info: "Sass Marketing landing",
    href: "/sass-marketing",
  },
  {
    title: "CRM Home",
    info: "Customer Relationshif",
    href: "/crm-home",
  },
  {
    title: "Help Desk",
    info: "Help desk home",
    href: "/help-desk",
  },
  {
    title: "Digital Agency",
    info: "Digital Agency home",
    href: "/digital-agency",
  },

  {
    title: "Software Company",
    info: "Software Company Home",
    href: "/software-company",
  },
  {
    title: "Creative Agency",
    info: "Creative Agency Landing One",
    href: "/creative-agency-one",
  },
  {
    title: "Digital Marketing Agency",
    info: "Digital Marketing Agency Landing",
    href: "/digital-marketing-agency",
  },
  {
    title: "Design Agency",
    info: "Disign Agency Home One",
    href: "/design-agency-one",
  },
  {
    title: "Design Agency",
    info: "Disign Agency Home Two",
    href: "/design-agency-two",
  },
  {
    title: "Agency Home",
    info: "Agency Home new",
    href: "/agency-home-new",
  },
  {
    title: "Creative Agency Two",
    info: "Creative Agency Landing Two",
    href: "/creative-agency-two",
  },
  {
    title: "Risk Managment",
    info: "Risk Managment home",
    href: "/risk-managment",
  },
  {
    title: "It Company",
    info: "It Company Home",
    href: "/it-company",
  },
  {
    title: "Ai Home",
    info: "Ai Company Home",
    href: "/ai-company",
  },
  {
    title: "Ai Content Generator",
    info: "Sass Content Generator",
    href: "/ai-content-generator",
  },
  {
    title: "Ai SAAS Content Generator",
    info: "Ai SAAS Content Marketing",
    href: "/ai-content-marketing",
  },
  {
    title: "Ai Image Generation SAAS",
    info: "Ai Image Generation SAAS",
    href: "/ai-image-generation",
  },
  {
    title: "Payment Gateway",
    info: "Payment Gateway New",
    href: "/payment-gateway-new",
  },
  {
    title: "Corporate Agency",
    info: "Corporate Agency",
    href: "/corporate-agency",
  },
  {
    title: "Creative Agency",
    info: "Creative Agency New",
    href: "/creative-agency-new",
  },
  {
    title: "Marketplace Agency",
    info: "Marketplace Agency",
    href: "/marketplace-agency",
  },
  {
    title: "Job Finder",
    info: "Job Finder New",
    href: "/job-finder",
  },
  {
    title: "Job Finder 2",
    info: "Job Finder New",
    href: "/job-finder-2",
  },
  {
    title: "Charity",
    info: "Charity home",
    href: "/charity",
  },
];

const TestimonialData = [
  {
    authorImg: "/img/testimonial/1.jpg",
    authorName: "Mr.Rupan Oberoi",
    authorTitle: "Founder and CEO at Amaara Herbs",
    quoteTitle: "The Best Template You Got to Have it!",
    authorQuote:
      "Globally network long-term high-impact schemas vis-a-vis distinctive e-commerce cross-media than ethical",
  },
  {
    id: 2,
    authorImg: "/img/testimonial/2.jpg",
    authorName: "Joan Dho",
    authorTitle: "Founder and CTO",
    quoteTitle: "Best Template for SAAS Company!",
    authorQuote:
      "Dynamically create innovative core competencies with effective best practices promote innovative infrastructures.",
  },
  {
    id: 3,
    authorImg: "/img/testimonial/3.jpg",
    authorName: "Ranu Mondal",
    authorTitle: "Lead Developer",
    quoteTitle: "It is undeniably good!",
    authorQuote:
      "Rapidiously supply client-centric e-markets and maintainable processes progressively engineer",
  },
  {
    id: 4,
    authorImg: "/img/testimonial/4.jpg",
    authorName: "Mr.Rupan Oberoi",
    authorTitle: "Founder and CEO at Amaara Herbs",
    quoteTitle: "The Best Template You Got to Have it!",
    authorQuote:
      "Globally network long-term high-impact schemas vis-a-vis distinctive e-commerce cross-media than ethical",
  },
];

const blogContent = [
  {
    title: "",
    subtitle: "",
    image: "/img/blog/blog-1.jpg",
  },
];
export {
  navCompanyLinks,
  navCompanyPage,
  offcanvasMenuData,
  TestimonialData,
  navServiceLinks,
  navInsightsLinks,
  navProductsLinks,
  blogContent,
};
