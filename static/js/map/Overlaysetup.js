import PlayerOverlay from './overlays/PlayerOverlay.js';
import PoiOverlay from './overlays/PoiOverlay.js';
import ShopOverlay from './overlays/ShopOverlay.js';
import LabelOverlay from './overlays/LabelOverlay.js';
import TrainlineOverlay from './overlays/TrainlineOverlay.js';
import TravelnetOverlay from './overlays/TravelnetOverlay.js';
import BonesOverlay from './overlays/BonesOverlay.js';
import LcdOverlay from './overlays/LcdOverlay.js';
import DigitermOverlay from './overlays/DigitermOverlay.js';
import LuacontrollerOverlay from './overlays/LuacontrollerOverlay.js';
import TechnicAnchorOverlay from './overlays/TechnicAnchorOverlay.js';
import TechnicQuarryOverlay from './overlays/TechnicQuarryOverlay.js';
import TechnicSwitchOverlay from './overlays/TechnicSwitchOverlay.js';
import ProtectorOverlay from './overlays/ProtectorOverlay.js';
import XPProtectorOverlay from './overlays/XPProtectorOverlay.js';
import PrivProtectorOverlay from './overlays/PrivProtectorOverlay.js';
import MissionOverlay from './overlays/MissionOverlay.js';
import MinecartOverlay from './overlays/MinecartOverlay.js';
import ATMOverlay from './overlays/ATMOverlay.js';
import LocatorOverlay from './overlays/LocatorOverlay.js';
import BorderOverlay from './overlays/BorderOverlay.js';
import TrainOverlay from './overlays/TrainOverlay.js';
import TrainsignalOverlay from './overlays/TrainsignalOverlay.js';

export default function(cfg, map, overlays, wsChannel, layerMgr){

  function isDefault(key){
    return cfg.defaultoverlays.indexOf(key) >= 0;
  }

  if (cfg.mapobjects.mapserver_player) {
    overlays.Player = new PlayerOverlay(wsChannel, layerMgr);
    if (isDefault("mapserver_player")) {
      map.addLayer(overlays.Player);
    }
  }

  if (cfg.mapobjects.mapserver_poi) {
    overlays.POI = new PoiOverlay(wsChannel, layerMgr);
    if (isDefault("mapserver_poi")) {
      map.addLayer(overlays.POI);
    }
  }

  if (cfg.mapobjects.smartshop || cfg.mapobjects.fancyvend) {
    overlays.Shop = new ShopOverlay(wsChannel, layerMgr);
    if (isDefault("smartshop") || isDefault("fancyvend")) {
      map.addLayer(overlays.Shop);
    }
  }

  if (cfg.mapobjects.mapserver_label) {
    overlays.Label = new LabelOverlay(wsChannel, layerMgr);
    if (isDefault("mapserver_label")) {
      map.addLayer(overlays.Label);
    }
  }

  if (cfg.mapobjects.mapserver_trainline) {
    overlays.Trainlines = new TrainlineOverlay(wsChannel, layerMgr);
    if (isDefault("mapserver_trainline")) {
      map.addLayer(overlays.Trainlines);
    }
  }

  if (cfg.mapobjects.mapserver_border) {
    overlays.Border = new BorderOverlay(wsChannel, layerMgr);
    if (isDefault("mapserver_border")) {
      map.addLayer(overlays.Border);
    }
  }

  if (cfg.mapobjects.travelnet) {
    overlays.Travelnet = new TravelnetOverlay(wsChannel, layerMgr);
    if (isDefault("travelnet")) {
      map.addLayer(overlays.Travelnet);
    }
  }

  if (cfg.mapobjects.bones) {
    overlays.Bones = new BonesOverlay(wsChannel, layerMgr);
    if (isDefault("bones")) {
      map.addLayer(overlays.Bones);
    }
  }

  if (cfg.mapobjects.digilines) {
    overlays["Digilines LCD"] = new LcdOverlay(wsChannel, layerMgr);
    if (isDefault("digilines")) {
      map.addLayer(overlays["Digilines LCD"]);
    }
  }

  if (cfg.mapobjects.digiterms) {
    overlays.Digiterms = new DigitermOverlay(wsChannel, layerMgr);
    if (isDefault("digiterms")) {
      map.addLayer(overlays.Digiterms);
    }
  }

  if (cfg.mapobjects.luacontroller) {
    overlays["Lua Controller"] = new LuacontrollerOverlay(wsChannel, layerMgr);
    if (isDefault("luacontroller")) {
      map.addLayer(overlays["Lua Controller"]);
    }
  }

  if (cfg.mapobjects.technic_anchor) {
    overlays["Technic Anchor"] = new TechnicAnchorOverlay(wsChannel, layerMgr);
    if (isDefault("technic_anchor")) {
      map.addLayer(overlays["Technic Anchor"]);
    }
  }

  if (cfg.mapobjects.technic_quarry) {
    overlays["Technic Quarry"] = new TechnicQuarryOverlay(wsChannel, layerMgr);
    if (isDefault("technic_quarry")) {
      map.addLayer(overlays["Technic Quarry"]);
    }
  }

  if (cfg.mapobjects.technic_switch) {
    overlays["Technic Switching station"] = new TechnicSwitchOverlay(wsChannel, layerMgr);
    if (isDefault("technic_switch")) {
      map.addLayer(overlays["Technic Switching station"]);
    }
  }

  if (cfg.mapobjects.protector) {
    overlays.Protector = new ProtectorOverlay(wsChannel, layerMgr);
    if (isDefault("protector")) {
      map.addLayer(overlays.Protector);
    }
  }

  if (cfg.mapobjects.xpprotector) {
    overlays["XP Protector"] = new XPProtectorOverlay(wsChannel, layerMgr);
    if (isDefault("xpprotector")) {
      map.addLayer(overlays["XP Protector"]);
    }
  }

  if (cfg.mapobjects.privprotector) {
    overlays["Priv Protector"] = new PrivProtectorOverlay(wsChannel, layerMgr);
    if (isDefault("privprotector")) {
      map.addLayer(overlays["Priv Protector"]);
    }
  }

  if (cfg.mapobjects.mission) {
    overlays.Missions = new MissionOverlay(wsChannel, layerMgr);
    if (isDefault("mission")) {
      map.addLayer(overlays.Missions);
    }
  }

  if (cfg.mapobjects.train) {
    overlays.Trains = new TrainOverlay(wsChannel, layerMgr);

    if (isDefault("train")) {
      map.addLayer(overlays.Trains);
    }
  }

  if (cfg.mapobjects.trainsignal) {
    overlays.Trainsignals = new TrainsignalOverlay(wsChannel, layerMgr);

    if (isDefault("trainsignal")) {
      map.addLayer(overlays.Trainsignals);
    }
  }

  if (cfg.mapobjects.minecart) {
    overlays.Minecart = new MinecartOverlay(wsChannel, layerMgr);
    if (isDefault("minecart")) {
      map.addLayer(overlays.Minecart);
    }
  }

  if (cfg.mapobjects.atm) {
    overlays.ATM = new ATMOverlay(wsChannel, layerMgr);
    if (isDefault("atm")) {
      map.addLayer(overlays.ATM);
    }
  }

  if (cfg.mapobjects.locator) {
    overlays.Locator = new LocatorOverlay(wsChannel, layerMgr);
    if (isDefault("locator")) {
      map.addLayer(overlays.Locator);
    }
  }

}