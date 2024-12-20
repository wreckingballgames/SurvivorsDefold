components {
  id: "basic_enemy"
  component: "/enemies/basic_enemy/basic_enemy.script"
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"enemies\"\n"
  "mask: \"players\"\n"
  "mask: \"enemies\"\n"
  "mask: \"player_abilities\"\n"
  "mask: \"enemy_detector\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 8.0\n"
  "}\n"
  "locked_rotation: true\n"
  ""
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"basic_enemy\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/graphics/sprites/sprites.atlas\"\n"
  "}\n"
  ""
  position {
    z: 1.0
  }
}
