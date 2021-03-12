import headImg0 from '../../public/profile/default0.png'
import headImg1 from '../../public/profile/default1.png'
import headImg2 from '../../public/profile/default2.png'
import headImg3 from '../../public/profile/default3.png'
import headImg4 from '../../public/profile/default4.png'
import headImg5 from '../../public/profile/default5.png'
import headImg6 from '../../public/profile/default6.png'
import headImg7 from '../../public/profile/default7.png'
import headImg8 from '../../public/profile/default8.png'
import headImg9 from '../../public/profile/default9.png'
import headImg10 from '../../public/profile/default10.png'
import headImg11 from '../../public/profile/default11.png'
import headImg12 from '../../public/profile/default12.png'
import headImg13 from '../../public/profile/default13.png'
import headImg14 from '../../public/profile/default14.png'
//================================================================//
import logoImg from '../../public/logo.png'
import teamFighter from '../../public/内卷系统.png'

let utils = {
    avatars: [
        headImg0,
        headImg1,
        headImg2,
        headImg3,
        headImg4,
        headImg5,
        headImg6,
        headImg7,
        headImg8,
        headImg9,
        headImg10,
        headImg11,
        headImg12,
        headImg13,
        headImg14,
    ],
    getColorFromId: (fromId) => {
        let color = "#FFF"
        switch (fromId.substr(fromId.length - 1, 1)) {
            case '0':
                color = "#ed5a65"
                break;
            case '1':
                color = "#b55176"
                break;
            case '2':
                color = "#8b2671"
                break;
            case '3':
                color = "#126bae"
                break;
            case '4':
                color = "#44889b"
                break;
            case '5':
                color = "#d6cb28"
                break;
            case '6':
                color = "#20894d"
                break;
            case '7':
                color = "#fcd337"
                break;
            case '8':
                color = "#ff9900"
                break;
            case '9':
                color = "#ac1f18"
                break;
            default:
                color = "#a55353"
                break;
        }
        return color
    },
    picLogo: logoImg,
    picTeamFighter: teamFighter,
}

export default utils;