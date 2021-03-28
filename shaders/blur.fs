#version 330

// Input vertex attributes (from vertex shader)
in vec2 fragTexCoord;
in vec4 fragColor;

// Input uniform values
uniform sampler2D texture0;
uniform vec4 colDiffuse;

// Output fragment color
out vec4 finalColor;

// NOTE: Add here your custom variables
uniform float size = 100.0;
uniform float quality = 3.0;
uniform float directions = 32.0;
uniform float threshold = 0.2;
uniform int colMapSize;
uniform sampler2D colMap;

void main()
{
    float pi = 6.28318530718; // Pi*2
   
    vec2 radius = size/textureSize(texture0, 0);
    
    // Pixel color
    vec4 color = texture(texture0, fragTexCoord);
    
    // Blur calculations
    for (float d=0.0; d<pi; d+=pi/directions) {
		for (float i=1.0/quality; i<=1.0; i+=1.0/quality) {
            color += texture(texture0, fragTexCoord+vec2(cos(d),sin(d))*radius*i);
        }
    }
    
    // Output to screen
    color /= quality * directions - 15.0;
    if (color.a > threshold) {
        float lowestDiff = 100000;
        vec4 diff;
        vec4 col;
        for (int i = 0; i < colMapSize; i++) {
            diff = color - texture(colMap, vec2(i, 0));
            if ((diff.r + diff.g + diff.b)/3 < lowestDiff) {
                col = texture(colMap, vec2(i, 0));
            }
        }
        finalColor = col;
    } else {
        finalColor = vec4(0);    
    }

    // Color set here
    //finalColor = texture(colMap, vec2(0, 0));
}
