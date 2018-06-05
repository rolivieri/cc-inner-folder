import os, shutil, sys
from distutils.dir_util import copy_tree

scriptsDir = os.path.dirname(os.path.realpath(__file__))
#print("Scripts folder: " + scriptsDir)
srcDir = scriptsDir + "/../src"
commonDir = srcDir + "/common"

if not os.path.isdir(commonDir) and not os.path.exists(commonDir):
    print("Something is wrong; the common folder does not exist: " + commonDir)
    sys.exit()

for root, dirs, files in os.walk(srcDir):
    for name in dirs:
        if name.endswith("CC"):
            targetDir = os.path.join(root, name) + "/common"
            print("targetDir: " + targetDir)
            shutil.rmtree(targetDir, True)
            print("About to copy folder '" + commonDir + "' to '" + targetDir + "'")
            copy_tree(commonDir, targetDir)
            print("Completed copying common folder successfully!")
            
