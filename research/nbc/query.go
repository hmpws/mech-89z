package main

const query = `
query bonanzaPage(
  $id: ID
  $name: String!
  $type: EntityPageType!
  $nationalBroadcastType: String
  $userId: String!
  $platform: SupportedPlatforms!
  $device: String
  $profile: JSON
  $ld: Boolean
  $oneApp: Boolean
  $timeZone: String
  $deepLinkHandle: String
  $app: NBCUBrands!
  $nbcAffiliateName: String
  $telemundoAffiliateName: String
  $language: Languages
  $playlistMachineName: String
  $mpxGuid: String
  $authorized: Boolean
  $minimumTiles: Int
  $endCardMpxGuid: String
  $endCardTagLine: String
  $isDayZero: Boolean
) {
  bonanzaPage(
    id: $id
    name: $name
    type: $type
    nationalBroadcastType: $nationalBroadcastType
    userId: $userId
    platform: $platform
    device: $device
    profile: $profile
    ld: $ld
    oneApp: $oneApp
    timeZone: $timeZone
    deepLinkHandle: $deepLinkHandle
    app: $app
    nbcAffiliateName: $nbcAffiliateName
    telemundoAffiliateName: $telemundoAffiliateName
    language: $language
    playlistMachineName: $playlistMachineName
    mpxGuid: $mpxGuid
    authorized: $authorized
    minimumTiles: $minimumTiles
    endCardMpxGuid: $endCardMpxGuid
    endCardTagLine: $endCardTagLine
    isDayZero: $isDayZero
  ) {
    id
    pageType
    name
    metadata {
      __typename
      ...videoPageData
      ...seriesHomepagePageData
      ...movieHomepagePageData
      ...brandLandingPageMetadata
      ...titleV2HomepagePageData
    }
    analytics {
      ...brandLandingPageAnalyticsAttributes
      ...pageAnalyticsAttributes
      ...videoPageAnalyticsAttributes
      ...titlePageAnalyticsAttributes
      ...titleV2PageAnalyticsAttributes
    }
    data {
      featured {
        ...slideshow
        ...hero
        ...ctaHero
      }
      notification {
        ...notification
      }
    }
  }
}

fragment image on Image {
  path
  width
  height
}
fragment component on Component {
  component
  meta
  treatments
}

fragment section on Section {
  logicName
  deepLinkHandle
}

fragment overlayData on OverlayData {
  content {
    ...descriptionSection
  }
}
fragment overlay on Overlay {
  ...component
  data {
    ...overlayData
  }
}
fragment descriptionSection on DescriptionSection {
  ...component
  data {
    ...descriptionData
  }
}

fragment slideshow on Slideshow {
  ...component
  ...section
  data {
    ...slideList
  }
  analytics {
    itemsList
    machineName
    listTitle
  }
}
fragment hero on Hero {
  ...component
  ...section
  data {
    ...heroData
  }
}
fragment ctaHero on CTAHero {
  ...component
  ...section
  data {
    ...ctaHeroData
  }
}

fragment hypermediaLink on HypermediaLink {
  linkTitle
  rel
  request {
    location
    method
    body
    headers {
      name
      value
    }
  }
}
fragment ctaLink on CTALink {
  ...component
  data {
    ...ctaData
  }
  analytics {
    smartDynamicCTA
    smartTileLabel
    smartTileLogic
    editorialLogic
    smartTileScenario
    ctaTitle
    destinationType
    destination
    brand {
      title
    }
    series
    movie
    isMovie
    videoTitle
    locked
    programmingType
    seasonNumber
    episodeNumber
    mpxGuid
    duration
    isPlaylist
    playlistMachineName
    playlistTitle
    isLive
    sponsorName
    isSponsoredTitle
    isTrailer
    liveEntitlement
    isVote
    isSportVideo
    language
    league
    event
    sport
  }
}

fragment upcomingModal on UpcomingModal {
  ...component
  data {
    ...upcomingModalData
  }
  analytics {
    modalName
    modalType
    dismissText
    programTitle
    brand {
      title
    }
  }
}

fragment slide on Slide {
  ...component
  data {
    ...slideItem
  }
  analytics {
    entityTitle
    isSponsoredSlide
    sponsorName
    dynamicallyGenerated
    dynamicGenerationLogic
  }
}

fragment upcomingLiveSlide on UpcomingLiveSlide {
  ...component
  data {
    ...upcomingLiveSlideData
  }
  analytics {
    analyticsType
    ctaLiveTitle
    ctaUpcomingTitle
    ctaNotInPackageTitle
    isLiveCallout
    isSponsoredSlide
    sponsorName
    programType
    genre
    secondaryGenre
    listOfGenres
    title
    secondaryTitle
    liveEntitlement
    league
    sport
    videoBroadcast
    nielsenClientId
    nielsenChannel
    nielsenSfCode
    isOlympics
    adobeVideoResearchTitle
    brand {
      title
    }
  }
}

fragment videoTile on VideoTile {
  ...component
  data {
    ...videoItem
  }
  analytics {
    brand {
      title
    }
    series
    title
    programmingType
    episodeNumber
    seasonNumber
    mpxGuid
    locked
    duration
    movie
    genre
    sport
    league
    language
    event
  }
}

fragment smartTile on SmartTile {
  ...component
  data {
    ...smartTileData
  }
  analytics {
    smartTileLabel
    smartTileLogic
    editorialLogic
    smartTileScenario
  }
}
fragment ctaSmartTile on CTASmartTile {
  ...component
  data {
    ...ctaSmartTileData
  }
  analytics {
    brand {
      title
    }
    title
    programmingType
    episodeNumber
    seasonNumber
    mpxGuid
    locked
    duration
    movie
    series
    genre
    smartTileLabel
    smartTileLogic
    editorialLogic
    smartTileScenario
    sponsorName
    isSponsoredTitle
  }
}

fragment componentData on ComponentData {
  instanceID
}

fragment ctaData on CTAData {
  ...componentData
  color
  gradientStart
  gradientEnd
  text
  destinationType
  destination
  endCardMpxGuid
  endCardTagLine
  playlistMachineName
  playlistCount
  urlAlias
  isLive
  isPlaylist
  title
  secondaryTitle
  secondaryTitleTag
  isTrailer
}

fragment descriptionData on DescriptionData {
  ...componentData
  optionalTitle
  description
  image
  shortTitle
  gradientStart
}
fragment brandLandingPageAnalyticsAttributes on BrandLandingPageAnalyticsAttributes {
  brand {
    title
  }
}
fragment pageAnalyticsAttributes on PageAnalyticsAttributes {
  series
  favoritedSeries
  brand {
    title
  }
}
fragment titlePageAnalyticsAttributes on TitlePageAnalyticsAttributes {
  series
  movie
  favoritedSeries
  brand {
    title
  }
}
fragment titleV2PageAnalyticsAttributes on TitleV2PageAnalyticsAttributes {
  series
  movie
  favoritedSeries
  genre
  secondaryGenres
  isTitleHub
  titleHub
  isSponsoredTitle
  sponsorName
  hasTrailerCTA
  pageType
  brand {
    title
  }
  emptyStateErrorDescription
  isEmpty
  category
  seriesType
  dartTag
  v4ID
  listOfGenres
}
fragment videoPageAnalyticsAttributes on VideoPageAnalyticsAttributes {
  convivaAssetName
}
fragment videoPageData on VideoPageData {
  mpxAccountId
}
fragment brandLandingPageMetadata on BrandLandingPageMetadata {
  whiteBrandLogo
  colorBrandLogo
  brandDisplayTitle
  lightPrimaryColor
  darkPrimaryColor
  brandLandingHeadline
  brandLandingDescription
  brandLandingLogo
  brandLandingBackgroundImage
  brandLandingBackgroundPreview
}
fragment titleV2HomepagePageData on TitleV2HomepagePageData {
  pageType
  isEmpty
  socialMedia {
    name
    url
    handle
  }
  referenceUrl
  shortDescription
  shortTitle
  isCoppaCompliant
  schemaType
  titleArt {
    ...image
  }
  numberOfEpisodes
  numberOfSeasons
  brandDisplayTitle
}
fragment movieHomepagePageData on MovieHomepagePageData {
  gradientStart
  gradientEnd
  lightPrimaryColor
  darkPrimaryColor
  brandLightPrimaryColor
  brandDarkPrimaryColor
  genres
  category
  socialMedia {
    name
    url
    handle
  }
  dartTag
  referenceUrl
  description
  shortDescription
  shortTitle
  isCoppaCompliant
  schemaType
  v4ID
  titleArt {
    path
    width
    height
  }
  multiPlatformLargeImage
  multiPlatformSmallImage
  credits {
    personFirstName
    personLastName
    characterFirstName
    characterLastName
  }
  whiteBrandLogo
  colorBrandLogo
  brandDisplayTitle
  canonicalUrl
  titleLogo
}
fragment seriesHomepagePageData on SeriesHomepagePageData {
  gradientStart
  gradientEnd
  lightPrimaryColor
  darkPrimaryColor
  brandLightPrimaryColor
  brandDarkPrimaryColor
  genres
  category
  seriesType
  socialMedia {
    name
    url
    handle
  }
  dartTag
  referenceUrl
  availableSeasons
  description
  shortDescription
  shortTitle
  isCoppaCompliant
  schemaType
  v4ID
  titleArt {
    path
    width
    height
  }
  multiPlatformLargeImage
  multiPlatformSmallImage
  credits {
    personFirstName
    personLastName
    characterFirstName
    characterLastName
  }
  numberOfEpisodes
  numberOfSeasons
  whiteBrandLogo
  colorBrandLogo
  brandDisplayTitle
  canonicalUrl
  titleLogo
}
fragment heroData on HeroData {
  ...componentData
  title
  secondaryTitle
  image
  compactImage
  favoriteID
  favoriteInteraction {
    default {
      ...hypermediaLink
    }
    undo {
      ...hypermediaLink
    }
  }
  smartTile {
    ...smartTile
  }
}
fragment ctaHeroData on CTAHeroData {
  ...componentData
  title
  titleLogo {
    ...image
  }
  gradientStart
  gradientEnd
  description
  secondaryDescription
  heroImage: image {
    ...image
  }
  heroCompactImage: compactImage {
    ...image
  }
  favoriteInteraction {
    default {
      ...hypermediaLink
    }
    undo {
      ...hypermediaLink
    }
  }
  sponsorLogo {
    ...image
  }
  sponsorLogoAltText
  sponsorName
  brandDisplayTitle
  colorBrandLogo {
    ...image
  }
  whiteBrandLogo {
    ...image
  }
  smartTileCTA {
    ...ctaSmartTile
  }
  aboutOverlay {
    ...overlay
  }
  primaryCTA {
    ...ctaLink
  }
  secondaryCTA {
    ...ctaLink
  }
}

fragment smartTileData on SmartTileData {
  ...componentData
  label
  tile {
    ...videoTile
  }
}

fragment ctaSmartTileData on CTASmartTileData {
  mpxAccountId
}

fragment item on Item {
  v4ID
  title
  secondaryTitle
  tertiaryTitle
  description
  image
  gradientStart
  gradientEnd
  labelBadge
  lastModified
}

fragment videoItem on VideoItem {
  mpxAccountId
}

fragment upcomingModalData on UpcomingModalData {
  machineName
  title
  description
  ctaText
  dismissText
  lastMinuteModalLifespan
  countdownDayLabel
  countdownHourLabel
  countdownMinLabel
  startTime
  backgroundImage
  backgroundVideo
  resourceId
  channelId
  streamAccessName
}

fragment slideItem on SlideItem {
  ...componentData
  ...item
  titleColor
  secondaryTitleColor
  description
  descriptionColor
  compactImage
  videoTitle
  percentViewed
  episodeNumber
  seasonNumber
  seriesShortTitle
  programmingType
  portraitPreview
  landscapePreview
  titleLogo
  brandDisplayTitle
  whiteBrandLogo
  colorBrandLogo
  tuneIn
  rating
  locked
  airDate
  sponsorLogo
  sponsorLogoAltText
  sponsorName
  externalAdvertiserId
  ariaLabel
  playlistBadge
  cta {
    ...ctaLink
    ...smartTile
  }
}
fragment upcomingLiveSlideData on UpcomingLiveSlideData {
  ...componentData
  v4ID
  title
  secondaryTitle
  description
  image
  gradientStart
  gradientEnd
  lastModified
  liveTuneIn
  upcomingTuneIn
  liveBadge
  titleColor
  secondaryTitleColor
  descriptionColor
  compactImage
  landscapePreview
  titleLogo
  brandDisplayTitle
  whiteBrandLogo
  colorBrandLogo
  sponsorLogo
  sponsorLogoAltText
  sponsorName
  startTime
  endTime
  liveAriaLabel
  upcomingAriaLabel
  liveCtaColor
  upcomingCtaColor
  liveCtaText
  upcomingCtaText
  notInPackageCtaText
  resourceId
  channelId
  machineName
  streamAccessName
  directToLiveThreshold
  upcomingModal {
    ...upcomingModal
  }
  notification {
    ...notification
  }
}

fragment slideList on SlideList {
  ...componentData
  lastModified
  items {
    ...slide
    ...upcomingLiveSlide
  }
}

fragment notification on Notification {
  ...component
  data {
    ...componentData
    v4ID
    machineName
    headline
    headlineColor
    message
    messageColor
    logo
    logoAltText
    portraitImage
    landscapeImage
    cta {
      ...ctaLink
    }
    dismissText
  }
  analytics {
    entityTitle
    dismissText
    placement
  }
}
`
